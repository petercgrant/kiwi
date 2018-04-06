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
	b *bytes.Buffer
}

func NewByteBuffer() *ByteBuffer {
	bb := ByteBuffer{
		b: bytes.NewBuffer(make([]byte, 0, InitialCapacity)),
	}
	return &bb
}

func NewSharedByteBuffer(data []byte) *ByteBuffer {
	bb := ByteBuffer{
		b: bytes.NewBuffer(data),
	}
	return &bb
}

func (bb *ByteBuffer) ReadBool(result *bool) bool {
	var value byte
	if !bb.ReadByte(&value) {
		*result = false
		return false
	}

	*result = !(value == 0)
	return true
}

func (bb *ByteBuffer) ReadByte(result *byte) bool {
	var err error
	*result, err = bb.b.ReadByte()
	return err == nil
}

func (bb *ByteBuffer) ReadVarFloat(result *float32) bool {
	var first byte
	if !bb.ReadByte(&first) {
		return false
	}

	// Optimization: use a single byte to store zero
	if first == 0 {
		*result = 0
		return true
	}

	// Endian-independent 32-bit read
	if bb.b.Len() < 3 {
		*result = 0
		return false
	}

	rest := make([]byte, 3)
	if _, err := bb.b.Read(rest); err != nil {
		return false
	}
	bits := uint32(first) | (uint32(rest[0]) << 8) | (uint32(rest[1]) << 16) | (uint32(rest[2]) << 24)

	// Move the exponent back into place
	bits = (bits << 23) | (bits >> 9)

	// Reinterpret as a floating-point number
	*result = math.Float32frombits(bits)
	return true
}
func (bb *ByteBuffer) ReadVarUint(result *uint32) bool {
	var shift, byte uint8
	value := uint32(0)
	for more := true; more; more = byte&128 != 0 && shift < 35 {
		if !bb.ReadByte(&byte) {
			*result = 0
			return false
		}
		value |= (uint32(byte) & 127) << shift
		shift += 7
	}
	*result = value
	return true
}
func (bb *ByteBuffer) ReadVarInt(result *int32) bool {
	var value uint32
	if !bb.ReadVarUint(&value) {
		*result = 0
		return false
	}

	if value&1 != 0 {
		*result = int32(^(value >> 1))
	} else {
		*result = int32(value >> 1)
	}
	return true
}
func (bb *ByteBuffer) ReadString(result *string) bool {
	var err error
	*result, err = bb.b.ReadString(0)
	return err == nil
}
func (bb *ByteBuffer) WriteBool(value bool) {
	if value {
		bb.WriteByte(1)
	} else {
		bb.WriteByte(0)
	}
}
func (bb *ByteBuffer) WriteByte(value byte) {
	bb.b.Grow(1)
	bb.b.WriteByte(value)
}
func (bb *ByteBuffer) WriteBytes(value []byte) {
	bb.b.Grow(len(value))
	bb.b.Write(value)
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
	var definitionCount uint32
	bs.definitions = nil

	if !bb.ReadVarUint(&definitionCount) {
		return false
	}

	bs.definitions = make([]Definition, definitionCount)
	for j := range bs.definitions {
		definition := &bs.definitions[j]
		var fieldCount uint32
		var kind byte
		if !bb.ReadString(&definition.Name) ||
			!bb.ReadByte(&kind) ||
			!bb.ReadVarUint(&fieldCount) ||
			!(kind != byte(KindEnum) && kind != byte(KindStruct) && kind != byte(KindMessage)) {
			return false
		}
		definition.Kind = Kind(kind)

		definition.Fields = make([]Field, fieldCount)

		for k := range definition.Fields {
			field := &definition.Fields[k]
			var fieldType int32
			if !bb.ReadString(&field.Name) ||
				!bb.ReadVarInt(&fieldType) ||
				!bb.ReadBool(&field.IsArray) ||
				!bb.ReadVarUint(&field.Value) ||
				fieldType < int32(FieldTypeString) ||
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

	if field.IsArray && !bb.ReadVarUint(&count) {
		return false
	}

	for ; count > 0; count-- {
		switch field.Type {
		case FieldTypeBool:
			fallthrough
		case FieldTypeByte:
			var dummy byte
			if !bb.ReadByte(&dummy) {
				return false
			}

		case FieldTypeInt:
			fallthrough
		case FieldTypeUint:
			var dummy uint32
			if !bb.ReadVarUint(&dummy) {
				return false
			}

		case FieldTypeFloat:
			var dummy float32
			if !bb.ReadVarFloat(&dummy) {
				return false
			}

		case FieldTypeString:
			var dummy byte
			for more := true; more; more = dummy != 0 {
				if !bb.ReadByte(&dummy) {
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
				var dummy uint32
				if !bb.ReadVarUint(&dummy) {
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
				var id uint32
				for true {
					if !bb.ReadVarUint(&id) {
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
