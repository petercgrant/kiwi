package kiwi

import (
	"bytes"
	"math"
)

const InitialCapacity = 256

const (
	FieldTypeBool FieldType = -iota - 1
	FieldTypeByte
	FieldTypeInt
	FieldTypeUint
	FieldTypeFloat
	FieldTypeString
)

const (
	KindEnum Kind = iota
	KindStruct
	KindMessage
)

type FieldType int32

type Kind byte

type ByteBuffer struct {
	*bytes.Buffer
}

func NewByteBuffer() *ByteBuffer {
	bb := ByteBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, InitialCapacity)),
	}
	return &bb
}

func NewSharedByteBuffer(data []byte) *ByteBuffer {
	bb := ByteBuffer{
		Buffer: bytes.NewBuffer(data),
	}
	return &bb
}

func (bb *ByteBuffer) ReadBool() (bool, bool) {
	value, ok := bb.ReadByte()
	if !ok {
		return false, false
	}

	return !(value == 0), true
}

func (bb *ByteBuffer) ReadByte() (byte, bool) {
	result, err := bb.Buffer.ReadByte()
	return result, err == nil
}

func (bb *ByteBuffer) ReadVarFloat() (float32, bool) {
	first, ok := bb.ReadByte()
	if !ok {
		return 0, false
	}

	// Optimization: use a single byte to store zero
	if first == 0 {
		return 0, true
	}

	// Endian-independent 32-bit read
	if bb.Buffer.Len() < 3 {
		return 0, false
	}

	rest := make([]byte, 3)
	if _, err := bb.Buffer.Read(rest); err != nil {
		return 0, false
	}
	bits := uint32(first) | (uint32(rest[0]) << 8) | (uint32(rest[1]) << 16) | (uint32(rest[2]) << 24)

	// Move the exponent back into place
	bits = (bits << 23) | (bits >> 9)

	// Reinterpret as a floating-point number
	return math.Float32frombits(bits), true
}
func (bb *ByteBuffer) ReadVarUint() (uint32, bool) {
	var shift, byte uint8
	var ok bool
	value := uint32(0)
	for more := true; more; more = byte&128 != 0 && shift < 35 {
		byte, ok = bb.ReadByte()
		if !ok {
			return 0, false
		}
		value |= (uint32(byte) & 127) << shift
		shift += 7
	}
	return value, true
}
func (bb *ByteBuffer) ReadVarInt() (int32, bool) {
	value, ok := bb.ReadVarUint()
	if !ok {
		return 0, false
	}

	var result int32
	if value&1 != 0 {
		result = int32(^(value >> 1))
	} else {
		result = int32(value >> 1)
	}
	return result, true
}
func (bb *ByteBuffer) ReadString() (string, bool) {
	result, err := bb.Buffer.ReadString(0)
	if err == nil && len(result) == 0 {
		return "", false
	}
	return result[:len(result)-1], err == nil
}
func (bb *ByteBuffer) WriteBool(value bool) {
	if value {
		bb.WriteByte(1)
	} else {
		bb.WriteByte(0)
	}
}
func (bb *ByteBuffer) WriteByte(value byte) {
	bb.Buffer.Grow(1)
	bb.Buffer.WriteByte(value)
}
func (bb *ByteBuffer) WriteBytes(value []byte) {
	bb.Buffer.Grow(len(value))
	bb.Buffer.Write(value)
}
func (bb *ByteBuffer) WriteVarFloat(value float32) {
	// Reinterpret as an integer
	bits := math.Float32bits(value)

	// Move the exponent to the first 8 bits
	bits = (bits >> 23) | (bits << 9)

	// Optimization: use a single byte to store zero and denormals (check for an exponent of 0)
	if bits&255 == 0 {
		bb.WriteByte(0)
		return
	}

	// Endian-independent 32-bit write
	data := []byte{byte(bits & 255), byte((bits >> 8) & 255), byte((bits >> 16) & 255), byte((bits >> 24) & 255)}
	bb.WriteBytes(data)
}
func (bb *ByteBuffer) WriteVarUint(value uint32) {
	for more := true; more; more = value != 0 {
		byte := byte(value & 127)
		value >>= 7
		if value == 0 {
			bb.WriteByte(byte)
		} else {
			bb.WriteByte(byte | 128)
		}
	}
}
func (bb *ByteBuffer) WriteVarInt(value int32) {
	bb.WriteVarUint((uint32(value) << 1) ^ (uint32(value) >> 31))
}
func (bb *ByteBuffer) WriteString(value string) {
	bb.WriteBytes([]byte(value))
	bb.WriteByte(0)
}

type BinarySchema struct {
	definitions []Definition
}

func (bs *BinarySchema) Parse(bb *ByteBuffer) bool {
	bs.definitions = nil

	definitionCount, ok := bb.ReadVarUint()
	if !ok {
		return false
	}

	bs.definitions = make([]Definition, definitionCount)
	for j := range bs.definitions {
		definition := &bs.definitions[j]
		definition.Name, ok = bb.ReadString()
		if !ok {
			return false
		}
		kind, ok := bb.ReadByte()
		if !ok {
			return false
		}
		fieldCount, ok := bb.ReadVarUint()
		if !ok {
			return false
		}

		if !(kind != byte(KindEnum) && kind != byte(KindStruct) && kind != byte(KindMessage)) {
			return false
		}
		definition.Kind = Kind(kind)

		definition.Fields = make([]Field, fieldCount)

		for k := range definition.Fields {
			field := &definition.Fields[k]
			var fieldType int32
			field.Name, ok = bb.ReadString()
			if !ok {
				return false
			}
			fieldType, ok := bb.ReadVarInt()
			if !ok {
				return false
			}
			field.IsArray, ok = bb.ReadBool()
			if !ok {
				return false
			}
			field.Value, ok = bb.ReadVarUint()
			if !ok {
				return false
			}

			if fieldType < int32(FieldTypeString) ||
				fieldType >= int32(definitionCount) {
				return false
			}
			field.Type = FieldType(fieldType)
		}
	}
	return true
}
func (bs *BinarySchema) FindDefinition(definition string, index *uint32) bool {
	for i := range bs.definitions {
		item := &bs.definitions[i]
		if item.Name == definition {
			*index = uint32(i)
			return true
		}
	}

	// Ignore fields we're looking for in an old schema
	*index = 0
	return false
}
func (bs *BinarySchema) SkipField(bb *ByteBuffer, definition uint32, field uint32) bool {
	if definition < uint32(len(bs.definitions)) {
		for j := range bs.definitions[definition].Fields {
			item := &bs.definitions[definition].Fields[j]
			if item.Value == field {
				return bs.skipField(bb, item)
			}
		}
	}

	return false
}

type Field struct {
	Name    string
	Type    FieldType
	IsArray bool
	Value   uint32
}

type Definition struct {
	Name   string
	Kind   Kind
	Fields []Field
}

func (bs *BinarySchema) skipField(bb *ByteBuffer, field *Field) bool {
	count := uint32(1)
	if field.IsArray {
		var ok bool
		if count, ok = bb.ReadVarUint(); !ok {
			return false
		}
	}

	for ; count > 0; count-- {
		switch field.Type {
		case FieldTypeBool:
			fallthrough
		case FieldTypeByte:
			if _, ok := bb.ReadByte(); !ok {
				return false
			}
		case FieldTypeInt:
			fallthrough
		case FieldTypeUint:
			if _, ok := bb.ReadVarUint(); !ok {
				return false
			}
		case FieldTypeFloat:
			if _, ok := bb.ReadVarFloat(); !ok {
				return false
			}
		case FieldTypeString:
			var dummy byte
			for more := true; more; more = dummy != 0 {
				var ok bool
				dummy, ok = bb.ReadByte()
				if !ok {
					return false
				}
			}
		default:
			if !(field.Type >= 0 && int32(field.Type) < int32(len(bs.definitions))) {
				return false
			}
			definition := bs.definitions[field.Type]

			switch definition.Kind {
			case KindEnum:
				if _, ok := bb.ReadVarUint(); !ok {
					return false
				}
			case KindStruct:
				for j := range definition.Fields {
					item := &definition.Fields[j]
					if !bs.skipField(bb, item) {
						return false
					}
				}

			case KindMessage:
				for true {
					id, ok := bb.ReadVarUint()
					if !ok {
						return false
					}
					if id == 0 {
						break
					}
					if !bs.SkipField(bb, uint32(field.Type), id) {
						return false
					}
				}
			default:
				return false
			}
		}
	}
	return true
}
