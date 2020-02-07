package kiwi

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/go-errors/errors"
	"github.com/picmonkey/go-spew/spew"
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
	n := len(data)
	if n == 0 {
		n = cap(data)
		if n == 0 {
			n = InitialCapacity
		}
	}
	bb := ByteBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, n)),
	}
	bb.Buffer.Write(data)
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
		if byte == 0 && value&127 > 0 {
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
	bb.WriteVarUint((uint32(int32(value)&0x7fffffff) << 1) ^ (uint32(int32(value) >> 31)))
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

type K interface{}
type V interface{}

type node struct {
	prev *node
	next *node
	k    K
	v    V
}

func (n *node) key() K {
	if n == nil {
		return nil
	}
	return n.k
}
func (n *node) value() V {
	if n == nil {
		return nil
	}
	return n.v
}

type linkedList struct {
	head *node
	tail *node
}

func (ll *linkedList) append(n *node) {
	n.prev = ll.tail
	n.next = nil
	if ll.tail == nil {
		ll.head = n
	} else {
		ll.tail.next = n
	}
	ll.tail = n
}
func (ll *linkedList) delete(n *node) {
	if prev := n.prev; prev != nil {
		prev.next = n.next
	}
	if next := n.next; next != nil {
		next.prev = n.prev
	}
	if n == ll.head {
		ll.head = n.next
	}
	if n == ll.tail {
		ll.tail = n.prev
	}
}
func NewLinkedMap(cap int) LinkedMap {
	return LinkedMap{m: make(map[K]*node, cap)}
}

type LinkedMap struct {
	m map[K]*node
	linkedList
}

func (m *LinkedMap) Set(k K, v V) {
	n, ok := m.m[k]
	if ok {
		n.v = v
	} else {
		n = &node{
			k: k,
			v: v,
		}
		m.append(n)
		m.m[k] = n
	}
}
func (m *LinkedMap) Get(k K) (V, bool) {
	n, ok := m.m[k]
	return n.value(), ok
}
func (m *LinkedMap) Delete(k K) {
	n, ok := m.m[k]
	if !ok {
		return
	}
	m.delete(n)
	delete(m.m, k)
}
func (m *LinkedMap) Len() int {
	return len(m.m)
}
func (m *LinkedMap) Print(printer *Printer) bool {
	printer.StartObject()
	for it := m.Iter(); it.HasNext(); {
		key, val := it.Next()
		printer.Print("key", key)
		printer.Print("val", val)
	}
	printer.EndObject()
	return true
}

type Iterator struct {
	cur *node
}

func (i *Iterator) HasNext() bool {
	return i.cur != nil
}
func (i *Iterator) Next() (K, V) {
	k, v := i.cur.key(), i.cur.value()
	if i.cur != nil {
		i.cur = i.cur.next
	}
	return k, v
}
func (m *LinkedMap) Iter() Iterator {
	return Iterator{m.head}
}

type Printer struct {
	parent *Printer
	seen   map[interface{}]bool
	indent int
	w      io.Writer
}

type Print interface {
	Print(p *Printer) bool
}

func (p *Printer) Field(name string) {
	if name == "" {
		fmt.Fprint(p.w, strings.Repeat("  ", p.indent))
	} else {
		fmt.Fprint(p.w, strings.Repeat("  ", p.indent), name, ": ")
	}
}

func (p *Printer) StartArray() {
	fmt.Fprintln(p.w, "[")
	p.indent++
}

func (p *Printer) EndArray() {
	p.indent--
	indent := strings.Repeat("  ", p.indent)
	fmt.Fprintf(p.w, "%s]\n", indent)
}

func (p *Printer) StartObject() {
	fmt.Fprintln(p.w, "{")
	p.indent++
}

func (p *Printer) EndObject() {
	p.indent--
	indent := strings.Repeat("  ", p.indent)
	fmt.Fprintf(p.w, "%s}\n", indent)
}

func (p *Printer) Print(name string, val interface{}) bool {
	if p == nil || p.seen == nil {
		return false
	}
	if val == nil {
		return true
	}
	if p.hasSeen(val) {
		return p.AlreadySeen(name, val)
	}
	p.Field(name)
	//p.seen[val] = true
	if v, ok := val.(Print); ok {
		return v.Print(p)
	}
	spew.Fdump(p.w, val)
	return true
}

func (p *Printer) hasSeen(val interface{}) bool {
	if p == nil {
		return false
	}
	if p.seen != nil {
		if p.seen[val] {
			return true
		}
	}
	return p.parent.hasSeen(val)
}

func (p *Printer) With(val interface{}) *Printer {
	return &Printer{
		parent: p,
		seen:   map[interface{}]bool{},
		indent: p.indent,
		w:      p.w,
	}
}

func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		seen:   map[interface{}]bool{},
		indent: 0,
		w:      w,
	}
}

func (p *Printer) AlreadySeen(name, val interface{}) bool {
	if p == nil || p.seen == nil {
		return false
	}
	p.emit("%s: (already seen)", name)
	return true
}

func (p *Printer) emit(format string, args ...interface{}) (int, error) {
	indent := strings.Repeat("  ", p.indent)
	n, err := fmt.Fprintf(p.w, indent+format+"\n", args...)
	if err != nil {
		return n, errors.Wrap(err, 0)
	}
	return n, nil
}
