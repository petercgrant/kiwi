package test

import "mnk.ee/kiwi"

type BinarySchema struct {
	schema kiwi.BinarySchema
	indexEnumMessage uint32
	indexBoolMessage uint32
	indexByteMessage uint32
	indexIntMessage uint32
	indexUintMessage uint32
	indexFloatMessage uint32
	indexFloat32ArrayMessage uint32
	indexStringMessage uint32
	indexCompoundMessage uint32
	indexMapMessage uint32
	indexNestedMessage uint32
	indexEnumArrayMessage uint32
	indexBoolArrayMessage uint32
	indexByteArrayMessage uint32
	indexIntArrayMessage uint32
	indexUintArrayMessage uint32
	indexFloatArrayMessage uint32
	indexStringArrayMessage uint32
	indexCompoundArrayMessage uint32
	indexRecursiveMessage uint32
	indexNonDeprecatedMessage uint32
	indexDeprecatedMessage uint32
}

type Enum uint32
const (

	Enum_A Enum = 100
	Enum_B Enum = 200
)


type EnumStruct struct {
	_flags [1]uint32
	_data_x Enum
	_data_y []Enum
}

type BoolStruct struct {
	_flags [1]uint32
	_data_x bool
}

type ByteStruct struct {
	_flags [1]uint32
	_data_x byte
}

type IntStruct struct {
	_flags [1]uint32
	_data_x int32
}

type UintStruct struct {
	_flags [1]uint32
	_data_x uint32
}

type FloatStruct struct {
	_flags [1]uint32
	_data_x float32
}

type StringStruct struct {
	_flags [1]uint32
	_data_x string
}

type CompoundStruct struct {
	_flags [1]uint32
	_data_x uint32
	_data_y uint32
}

type NestedStruct struct {
	_flags [1]uint32
	_data_b *CompoundStruct
	_data_d kiwi.LinkedMap
	_data_a uint32
	_data_c uint32
}

type EnumMessage struct {
	_flags [1]uint32
	_data_x Enum
}

type BoolMessage struct {
	_flags [1]uint32
	_data_x bool
}

type ByteMessage struct {
	_flags [1]uint32
	_data_x byte
}

type IntMessage struct {
	_flags [1]uint32
	_data_x int32
}

type UintMessage struct {
	_flags [1]uint32
	_data_x uint32
}

type FloatMessage struct {
	_flags [1]uint32
	_data_x float32
}

type Float32ArrayMessage struct {
	_flags [1]uint32
	_data_x []float32
}

type StringMessage struct {
	_flags [1]uint32
	_data_x string
}

type CompoundMessage struct {
	_flags [1]uint32
	_data_x uint32
	_data_y uint32
}

type MapMessage struct {
	_flags [1]uint32
	_data_x kiwi.LinkedMap
}

type NestedMessage struct {
	_flags [1]uint32
	_data_b *CompoundMessage
	_data_d kiwi.LinkedMap
	_data_a uint32
	_data_c uint32
}

type EnumArrayStruct struct {
	_flags [1]uint32
	_data_x []Enum
}

type BoolArrayStruct struct {
	_flags [1]uint32
	_data_x []bool
}

type ByteArrayStruct struct {
	_flags [1]uint32
	_data_x []byte
}

type IntArrayStruct struct {
	_flags [1]uint32
	_data_x []int32
}

type UintArrayStruct struct {
	_flags [1]uint32
	_data_x []uint32
}

type FloatArrayStruct struct {
	_flags [1]uint32
	_data_x []float32
}

type StringArrayStruct struct {
	_flags [1]uint32
	_data_x []string
}

type CompoundArrayStruct struct {
	_flags [1]uint32
	_data_x []uint32
	_data_y []uint32
}

type EnumArrayMessage struct {
	_flags [1]uint32
	_data_x []Enum
}

type BoolArrayMessage struct {
	_flags [1]uint32
	_data_x []bool
}

type ByteArrayMessage struct {
	_flags [1]uint32
	_data_x []byte
}

type IntArrayMessage struct {
	_flags [1]uint32
	_data_x []int32
}

type UintArrayMessage struct {
	_flags [1]uint32
	_data_x []uint32
}

type FloatArrayMessage struct {
	_flags [1]uint32
	_data_x []float32
}

type StringArrayMessage struct {
	_flags [1]uint32
	_data_x []string
}

type CompoundArrayMessage struct {
	_flags [1]uint32
	_data_x []uint32
	_data_y []uint32
}

type RecursiveMessage struct {
	_flags [1]uint32
	_data_x *RecursiveMessage
}

type NonDeprecatedMessage struct {
	_flags [1]uint32
	_data_c []uint32
	_data_d []uint32
	_data_e *ByteStruct
	_data_f *ByteStruct
	_data_a uint32
	_data_b uint32
	_data_g uint32
}

type DeprecatedMessage struct {
	_flags [1]uint32
	_data_c []uint32
	_data_e *ByteStruct
	_data_a uint32
	_data_g uint32
}

type SortedStruct struct {
	_flags [1]uint32
	_data_f1 string
	_data_f2 string
	_data_a3 []bool
	_data_b3 []byte
	_data_c3 []int32
	_data_d3 []uint32
	_data_e3 []float32
	_data_f3 []string
	_data_c1 int32
	_data_d1 uint32
	_data_e1 float32
	_data_c2 int32
	_data_d2 uint32
	_data_e2 float32
	_data_a1 bool
	_data_b1 byte
	_data_a2 bool
	_data_b2 byte
}


func (bs *BinarySchema) Parse(bb *kiwi.ByteBuffer) bool {
	if !bs.schema.Parse(bb) {
		return false
	}

	bs.schema.FindDefinition("EnumMessage", &bs.indexEnumMessage)
	bs.schema.FindDefinition("BoolMessage", &bs.indexBoolMessage)
	bs.schema.FindDefinition("ByteMessage", &bs.indexByteMessage)
	bs.schema.FindDefinition("IntMessage", &bs.indexIntMessage)
	bs.schema.FindDefinition("UintMessage", &bs.indexUintMessage)
	bs.schema.FindDefinition("FloatMessage", &bs.indexFloatMessage)
	bs.schema.FindDefinition("Float32ArrayMessage", &bs.indexFloat32ArrayMessage)
	bs.schema.FindDefinition("StringMessage", &bs.indexStringMessage)
	bs.schema.FindDefinition("CompoundMessage", &bs.indexCompoundMessage)
	bs.schema.FindDefinition("MapMessage", &bs.indexMapMessage)
	bs.schema.FindDefinition("NestedMessage", &bs.indexNestedMessage)
	bs.schema.FindDefinition("EnumArrayMessage", &bs.indexEnumArrayMessage)
	bs.schema.FindDefinition("BoolArrayMessage", &bs.indexBoolArrayMessage)
	bs.schema.FindDefinition("ByteArrayMessage", &bs.indexByteArrayMessage)
	bs.schema.FindDefinition("IntArrayMessage", &bs.indexIntArrayMessage)
	bs.schema.FindDefinition("UintArrayMessage", &bs.indexUintArrayMessage)
	bs.schema.FindDefinition("FloatArrayMessage", &bs.indexFloatArrayMessage)
	bs.schema.FindDefinition("StringArrayMessage", &bs.indexStringArrayMessage)
	bs.schema.FindDefinition("CompoundArrayMessage", &bs.indexCompoundArrayMessage)
	bs.schema.FindDefinition("RecursiveMessage", &bs.indexRecursiveMessage)
	bs.schema.FindDefinition("NonDeprecatedMessage", &bs.indexNonDeprecatedMessage)
	bs.schema.FindDefinition("DeprecatedMessage", &bs.indexDeprecatedMessage)
	return true
}

func (bs *BinarySchema) SkipEnumMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexEnumMessage, id)
}

func (bs *BinarySchema) SkipBoolMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexBoolMessage, id)
}

func (bs *BinarySchema) SkipByteMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexByteMessage, id)
}

func (bs *BinarySchema) SkipIntMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexIntMessage, id)
}

func (bs *BinarySchema) SkipUintMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexUintMessage, id)
}

func (bs *BinarySchema) SkipFloatMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexFloatMessage, id)
}

func (bs *BinarySchema) SkipFloat32ArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexFloat32ArrayMessage, id)
}

func (bs *BinarySchema) SkipStringMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexStringMessage, id)
}

func (bs *BinarySchema) SkipCompoundMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexCompoundMessage, id)
}

func (bs *BinarySchema) SkipMapMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexMapMessage, id)
}

func (bs *BinarySchema) SkipNestedMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexNestedMessage, id)
}

func (bs *BinarySchema) SkipEnumArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexEnumArrayMessage, id)
}

func (bs *BinarySchema) SkipBoolArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexBoolArrayMessage, id)
}

func (bs *BinarySchema) SkipByteArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexByteArrayMessage, id)
}

func (bs *BinarySchema) SkipIntArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexIntArrayMessage, id)
}

func (bs *BinarySchema) SkipUintArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexUintArrayMessage, id)
}

func (bs *BinarySchema) SkipFloatArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexFloatArrayMessage, id)
}

func (bs *BinarySchema) SkipStringArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexStringArrayMessage, id)
}

func (bs *BinarySchema) SkipCompoundArrayMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexCompoundArrayMessage, id)
}

func (bs *BinarySchema) SkipRecursiveMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexRecursiveMessage, id)
}

func (bs *BinarySchema) SkipNonDeprecatedMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexNonDeprecatedMessage, id)
}

func (bs *BinarySchema) SkipDeprecatedMessageField(bb *kiwi.ByteBuffer, id uint32) bool {
	return bs.schema.SkipField(bb, bs.indexDeprecatedMessage, id)
}

func (p *EnumStruct) Getx() *Enum {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *EnumStruct) Setx(value Enum) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *EnumStruct) Gety() []Enum {
	if p._flags[0] & 2 > 0 {
		return p._data_y
	}
	return nil
}

func (p *EnumStruct) Sety(count uint32) []Enum {
	p._flags[0] |= 2
	p._data_y = make([]Enum, int(count))
	return p._data_y
}

func (p *EnumStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(p._data_x))
	if p.Gety() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_y)))
	for _, _it := range p._data_y{
		bb.WriteVarUint(uint32(_it))
	}
	return true;
}

func (p *EnumStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	if p._data_x, ok = func() (val Enum, ok bool) {
		var v uint32
		if v, ok = bb.ReadVarUint(); ok {
			val = Enum(v)
		}
		return
	}(); !ok {
		return false
	}
	p.Setx(p._data_x)
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	y := p.Sety(count)
	for j := range y {
		y[j], ok = func() (val Enum, ok bool) {
			var v uint32
			if v, ok = bb.ReadVarUint(); ok {
				val = Enum(v)
			}
			return
		}()
		if !ok {
			return false
		}
	}
	return true
}

func (p *EnumStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	if p.Gety() != nil {
		printer.Field("y")
		printer.StartArray()
		for _, _it := range p._data_y{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *BoolStruct) Getx() *bool {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *BoolStruct) Setx(value bool) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *BoolStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteBool(p._data_x)
	return true;
}

func (p *BoolStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadBool(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *BoolStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *ByteStruct) Getx() *byte {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *ByteStruct) Setx(value byte) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *ByteStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteByte(p._data_x)
	return true;
}

func (p *ByteStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadByte(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *ByteStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *IntStruct) Getx() *int32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *IntStruct) Setx(value int32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *IntStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarInt(p._data_x)
	return true;
}

func (p *IntStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadVarInt(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *IntStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *UintStruct) Getx() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *UintStruct) Setx(value uint32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *UintStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(p._data_x)
	return true;
}

func (p *UintStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *UintStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *FloatStruct) Getx() *float32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *FloatStruct) Setx(value float32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *FloatStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarFloat(p._data_x)
	return true;
}

func (p *FloatStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadVarFloat(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *FloatStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *StringStruct) Getx() *string {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *StringStruct) Setx(value string) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *StringStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteString(p._data_x)
	return true;
}

func (p *StringStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadString(); !ok {
		return false
	}
	p.Setx(p._data_x)
	return true
}

func (p *StringStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *CompoundStruct) Getx() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *CompoundStruct) Setx(value uint32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *CompoundStruct) Gety() *uint32 {
	if p._flags[0] & 2 > 0 {
		return &p._data_y
	}
	return nil
}

func (p *CompoundStruct) Sety(value uint32) {
	p._flags[0] |= 2
	p._data_y = value
}

func (p *CompoundStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(p._data_x)
	if p.Gety() == nil {
		return false
	}
	bb.WriteVarUint(p._data_y)
	return true;
}

func (p *CompoundStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	if p._data_x, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Setx(p._data_x)
	if p._data_y, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Sety(p._data_y)
	return true
}

func (p *CompoundStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	if p.Gety() != nil {
		printer.Print("y", p.Gety())
	}
	printer.EndObject()
	return true;
}

func (p *NestedStruct) Geta() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_a
	}
	return nil
}

func (p *NestedStruct) Seta(value uint32) {
	p._flags[0] |= 1
	p._data_a = value
}

func (p *NestedStruct) Getb() *CompoundStruct {
	if p == nil {
		return nil
	}
	return p._data_b
}

func (p *NestedStruct) Setb(value *CompoundStruct) {
	p._data_b = value;
}

func (p *NestedStruct) Getc() *uint32 {
	if p._flags[0] & 4 > 0 {
		return &p._data_c
	}
	return nil
}

func (p *NestedStruct) Setc(value uint32) {
	p._flags[0] |= 4
	p._data_c = value
}

func (p *NestedStruct) Getd() *kiwi.LinkedMap {
	if p._flags[0] & 8 > 0 {
		return &p._data_d
	}
	return nil
}

func (p *NestedStruct) Setd(value kiwi.LinkedMap) {
	p._flags[0] |= 8
	p._data_d = value
}

func (p *NestedStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Geta() == nil {
		return false
	}
	bb.WriteVarUint(p._data_a)
	if p.Getb() == nil {
		return false
	}
	if !p._data_b.Encode(bb) {
		return false
	}
	if p.Getc() == nil {
		return false
	}
	bb.WriteVarUint(p._data_c)
	if p.Getd() == nil {
		return false
	}
	bb.WriteVarUint(uint32(p._data_d.Len()))
	for it := p._data_d.Iter(); it.HasNext(); {
		kif, vif := it.Next()
		key := kif.(string)
		val := vif.(*CompoundStruct)
		bb.WriteString(key)
		if !val.Encode(bb) {
			return false
		}
	}
	return true;
}

func (p *NestedStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	if p._data_a, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Seta(p._data_a)
	if p._data_b, ok = func() (val *CompoundStruct, ok bool) {
		v := &CompoundStruct{}
		if ok = v.Decode(bb, schema); ok {
			val = v
		}
		return
	}(); !ok {
		return false
	}
	if p._data_c, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Setc(p._data_c)
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	p.Setd(kiwi.NewLinkedMap(int(count)))
	for j := uint32(0); j < count; j++ {
		key, ok := bb.ReadString()
		if !ok {
			return false
		}
		val, ok := func() (val *CompoundStruct, ok bool) {
			v := &CompoundStruct{}
			if ok = v.Decode(bb, schema); ok {
				val = v
			}
			return
		}()
		if !ok {
			return false
		}
		p._data_d.Set(key, val)
	}
	return true
}

func (p *NestedStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Geta() != nil {
		printer.Print("a", p.Geta())
	}
	if p.Getb() != nil {
		printer.Print("b", p.Getb())
	}
	if p.Getc() != nil {
		printer.Print("c", p.Getc())
	}
	if p.Getd() != nil {
		printer.Print("d", p.Getd())
	}
	printer.EndObject()
	return true;
}

func (p *EnumMessage) Getx() *Enum {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *EnumMessage) Setx(value Enum) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *EnumMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(p._data_x))
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *EnumMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = func() (val Enum, ok bool) {
					var v uint32
					if v, ok = bb.ReadVarUint(); ok {
						val = Enum(v)
					}
					return
				}(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipEnumMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *EnumMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *BoolMessage) Getx() *bool {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *BoolMessage) Setx(value bool) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *BoolMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteBool(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *BoolMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadBool(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipBoolMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *BoolMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *ByteMessage) Getx() *byte {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *ByteMessage) Setx(value byte) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *ByteMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteByte(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *ByteMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadByte(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipByteMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *ByteMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *IntMessage) Getx() *int32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *IntMessage) Setx(value int32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *IntMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarInt(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *IntMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadVarInt(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipIntMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *IntMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *UintMessage) Getx() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *UintMessage) Setx(value uint32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *UintMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *UintMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipUintMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *UintMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *FloatMessage) Getx() *float32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *FloatMessage) Setx(value float32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *FloatMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarFloat(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *FloatMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadVarFloat(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipFloatMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *FloatMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *Float32ArrayMessage) Getx() []float32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *Float32ArrayMessage) Setx(count uint32) []float32 {
	p._flags[0] |= 1
	p._data_x = make([]float32, int(count))
	return p._data_x
}

func (p *Float32ArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarFloat(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *Float32ArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadVarFloat()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipFloat32ArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *Float32ArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *StringMessage) Getx() *string {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *StringMessage) Setx(value string) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *StringMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteString(p._data_x)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *StringMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadString(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			default:
				if schema == nil || !schema.SkipStringMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *StringMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *CompoundMessage) Getx() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *CompoundMessage) Setx(value uint32) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *CompoundMessage) Gety() *uint32 {
	if p._flags[0] & 2 > 0 {
		return &p._data_y
	}
	return nil
}

func (p *CompoundMessage) Sety(value uint32) {
	p._flags[0] |= 2
	p._data_y = value
}

func (p *CompoundMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(p._data_x)
	}
	if p.Gety() != nil {
		bb.WriteVarUint(2)
		bb.WriteVarUint(p._data_y)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *CompoundMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setx(p._data_x)
				break
			case 2:
				if p._data_y, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Sety(p._data_y)
				break
			default:
				if schema == nil || !schema.SkipCompoundMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *CompoundMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	if p.Gety() != nil {
		printer.Print("y", p.Gety())
	}
	printer.EndObject()
	return true;
}

func (p *MapMessage) Getx() *kiwi.LinkedMap {
	if p._flags[0] & 1 > 0 {
		return &p._data_x
	}
	return nil
}

func (p *MapMessage) Setx(value kiwi.LinkedMap) {
	p._flags[0] |= 1
	p._data_x = value
}

func (p *MapMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(p._data_x.Len()))
		for it := p._data_x.Iter(); it.HasNext(); {
			kif, vif := it.Next()
			key := kif.(string)
			val := vif.(int32)
			bb.WriteString(key)
			bb.WriteVarInt(val)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *MapMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				p.Setx(kiwi.NewLinkedMap(int(count)))
				for j := uint32(0); j < count; j++ {
					key, ok := bb.ReadString()
					if !ok {
						return false
					}
					val, ok := bb.ReadVarInt()
					if !ok {
						return false
					}
					p._data_x.Set(key, val)
				}
				break
			default:
				if schema == nil || !schema.SkipMapMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *MapMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *NestedMessage) Geta() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_a
	}
	return nil
}

func (p *NestedMessage) Seta(value uint32) {
	p._flags[0] |= 1
	p._data_a = value
}

func (p *NestedMessage) Getb() *CompoundMessage {
	if p == nil {
		return nil
	}
	return p._data_b
}

func (p *NestedMessage) Setb(value *CompoundMessage) {
	p._data_b = value;
}

func (p *NestedMessage) Getc() *uint32 {
	if p._flags[0] & 4 > 0 {
		return &p._data_c
	}
	return nil
}

func (p *NestedMessage) Setc(value uint32) {
	p._flags[0] |= 4
	p._data_c = value
}

func (p *NestedMessage) Getd() *kiwi.LinkedMap {
	if p._flags[0] & 8 > 0 {
		return &p._data_d
	}
	return nil
}

func (p *NestedMessage) Setd(value kiwi.LinkedMap) {
	p._flags[0] |= 8
	p._data_d = value
}

func (p *NestedMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Geta() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(p._data_a)
	}
	if p.Getb() != nil {
		bb.WriteVarUint(2)
		if !p._data_b.Encode(bb) {
			return false
		}
	}
	if p.Getc() != nil {
		bb.WriteVarUint(3)
		bb.WriteVarUint(p._data_c)
	}
	if p.Getd() != nil {
		bb.WriteVarUint(4)
		bb.WriteVarUint(uint32(p._data_d.Len()))
		for it := p._data_d.Iter(); it.HasNext(); {
			kif, vif := it.Next()
			key := kif.(string)
			val := vif.(*CompoundMessage)
			bb.WriteString(key)
			if !val.Encode(bb) {
				return false
			}
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *NestedMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_a, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Seta(p._data_a)
				break
			case 2:
				if p._data_b, ok = func() (val *CompoundMessage, ok bool) {
					v := &CompoundMessage{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			case 3:
				if p._data_c, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setc(p._data_c)
				break
			case 4:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				p.Setd(kiwi.NewLinkedMap(int(count)))
				for j := uint32(0); j < count; j++ {
					key, ok := bb.ReadString()
					if !ok {
						return false
					}
					val, ok := func() (val *CompoundMessage, ok bool) {
						v := &CompoundMessage{}
						if ok = v.Decode(bb, schema); ok {
							val = v
						}
						return
					}()
					if !ok {
						return false
					}
					p._data_d.Set(key, val)
				}
				break
			default:
				if schema == nil || !schema.SkipNestedMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *NestedMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Geta() != nil {
		printer.Print("a", p.Geta())
	}
	if p.Getb() != nil {
		printer.Print("b", p.Getb())
	}
	if p.Getc() != nil {
		printer.Print("c", p.Getc())
	}
	if p.Getd() != nil {
		printer.Print("d", p.Getd())
	}
	printer.EndObject()
	return true;
}

func (p *EnumArrayStruct) Getx() []Enum {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *EnumArrayStruct) Setx(count uint32) []Enum {
	p._flags[0] |= 1
	p._data_x = make([]Enum, int(count))
	return p._data_x
}

func (p *EnumArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteVarUint(uint32(_it))
	}
	return true;
}

func (p *EnumArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = func() (val Enum, ok bool) {
			var v uint32
			if v, ok = bb.ReadVarUint(); ok {
				val = Enum(v)
			}
			return
		}()
		if !ok {
			return false
		}
	}
	return true
}

func (p *EnumArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *BoolArrayStruct) Getx() []bool {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *BoolArrayStruct) Setx(count uint32) []bool {
	p._flags[0] |= 1
	p._data_x = make([]bool, int(count))
	return p._data_x
}

func (p *BoolArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteBool(_it)
	}
	return true;
}

func (p *BoolArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadBool()
		if !ok {
			return false
		}
	}
	return true
}

func (p *BoolArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *ByteArrayStruct) Getx() []byte {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *ByteArrayStruct) Setx(count uint32) []byte {
	p._flags[0] |= 1
	p._data_x = make([]byte, int(count))
	return p._data_x
}

func (p *ByteArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteByte(_it)
	}
	return true;
}

func (p *ByteArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadByte()
		if !ok {
			return false
		}
	}
	return true
}

func (p *ByteArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *IntArrayStruct) Getx() []int32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *IntArrayStruct) Setx(count uint32) []int32 {
	p._flags[0] |= 1
	p._data_x = make([]int32, int(count))
	return p._data_x
}

func (p *IntArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteVarInt(_it)
	}
	return true;
}

func (p *IntArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadVarInt()
		if !ok {
			return false
		}
	}
	return true
}

func (p *IntArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *UintArrayStruct) Getx() []uint32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *UintArrayStruct) Setx(count uint32) []uint32 {
	p._flags[0] |= 1
	p._data_x = make([]uint32, int(count))
	return p._data_x
}

func (p *UintArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteVarUint(_it)
	}
	return true;
}

func (p *UintArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadVarUint()
		if !ok {
			return false
		}
	}
	return true
}

func (p *UintArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *FloatArrayStruct) Getx() []float32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *FloatArrayStruct) Setx(count uint32) []float32 {
	p._flags[0] |= 1
	p._data_x = make([]float32, int(count))
	return p._data_x
}

func (p *FloatArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteVarFloat(_it)
	}
	return true;
}

func (p *FloatArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadVarFloat()
		if !ok {
			return false
		}
	}
	return true
}

func (p *FloatArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *StringArrayStruct) Getx() []string {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *StringArrayStruct) Setx(count uint32) []string {
	p._flags[0] |= 1
	p._data_x = make([]string, int(count))
	return p._data_x
}

func (p *StringArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteString(_it)
	}
	return true;
}

func (p *StringArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadString()
		if !ok {
			return false
		}
	}
	return true
}

func (p *StringArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *CompoundArrayStruct) Getx() []uint32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *CompoundArrayStruct) Setx(count uint32) []uint32 {
	p._flags[0] |= 1
	p._data_x = make([]uint32, int(count))
	return p._data_x
}

func (p *CompoundArrayStruct) Gety() []uint32 {
	if p._flags[0] & 2 > 0 {
		return p._data_y
	}
	return nil
}

func (p *CompoundArrayStruct) Sety(count uint32) []uint32 {
	p._flags[0] |= 2
	p._data_y = make([]uint32, int(count))
	return p._data_y
}

func (p *CompoundArrayStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_x)))
	for _, _it := range p._data_x{
		bb.WriteVarUint(_it)
	}
	if p.Gety() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_y)))
	for _, _it := range p._data_y{
		bb.WriteVarUint(_it)
	}
	return true;
}

func (p *CompoundArrayStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	x := p.Setx(count)
	for j := range x {
		x[j], ok = bb.ReadVarUint()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	y := p.Sety(count)
	for j := range y {
		y[j], ok = bb.ReadVarUint()
		if !ok {
			return false
		}
	}
	return true
}

func (p *CompoundArrayStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Gety() != nil {
		printer.Field("y")
		printer.StartArray()
		for _, _it := range p._data_y{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *EnumArrayMessage) Getx() []Enum {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *EnumArrayMessage) Setx(count uint32) []Enum {
	p._flags[0] |= 1
	p._data_x = make([]Enum, int(count))
	return p._data_x
}

func (p *EnumArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarUint(uint32(_it))
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *EnumArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = func() (val Enum, ok bool) {
						var v uint32
						if v, ok = bb.ReadVarUint(); ok {
							val = Enum(v)
						}
						return
					}()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipEnumArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *EnumArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *BoolArrayMessage) Getx() []bool {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *BoolArrayMessage) Setx(count uint32) []bool {
	p._flags[0] |= 1
	p._data_x = make([]bool, int(count))
	return p._data_x
}

func (p *BoolArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteBool(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *BoolArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadBool()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipBoolArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *BoolArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *ByteArrayMessage) Getx() []byte {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *ByteArrayMessage) Setx(count uint32) []byte {
	p._flags[0] |= 1
	p._data_x = make([]byte, int(count))
	return p._data_x
}

func (p *ByteArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteByte(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *ByteArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadByte()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipByteArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *ByteArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *IntArrayMessage) Getx() []int32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *IntArrayMessage) Setx(count uint32) []int32 {
	p._flags[0] |= 1
	p._data_x = make([]int32, int(count))
	return p._data_x
}

func (p *IntArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarInt(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *IntArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadVarInt()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipIntArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *IntArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *UintArrayMessage) Getx() []uint32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *UintArrayMessage) Setx(count uint32) []uint32 {
	p._flags[0] |= 1
	p._data_x = make([]uint32, int(count))
	return p._data_x
}

func (p *UintArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarUint(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *UintArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipUintArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *UintArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *FloatArrayMessage) Getx() []float32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *FloatArrayMessage) Setx(count uint32) []float32 {
	p._flags[0] |= 1
	p._data_x = make([]float32, int(count))
	return p._data_x
}

func (p *FloatArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarFloat(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *FloatArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadVarFloat()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipFloatArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *FloatArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *StringArrayMessage) Getx() []string {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *StringArrayMessage) Setx(count uint32) []string {
	p._flags[0] |= 1
	p._data_x = make([]string, int(count))
	return p._data_x
}

func (p *StringArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteString(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *StringArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadString()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipStringArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *StringArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *CompoundArrayMessage) Getx() []uint32 {
	if p._flags[0] & 1 > 0 {
		return p._data_x
	}
	return nil
}

func (p *CompoundArrayMessage) Setx(count uint32) []uint32 {
	p._flags[0] |= 1
	p._data_x = make([]uint32, int(count))
	return p._data_x
}

func (p *CompoundArrayMessage) Gety() []uint32 {
	if p._flags[0] & 2 > 0 {
		return p._data_y
	}
	return nil
}

func (p *CompoundArrayMessage) Sety(count uint32) []uint32 {
	p._flags[0] |= 2
	p._data_y = make([]uint32, int(count))
	return p._data_y
}

func (p *CompoundArrayMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(uint32(len(p._data_x)))
		for _, _it := range p._data_x{
			bb.WriteVarUint(_it)
		}
	}
	if p.Gety() != nil {
		bb.WriteVarUint(2)
		bb.WriteVarUint(uint32(len(p._data_y)))
		for _, _it := range p._data_y{
			bb.WriteVarUint(_it)
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *CompoundArrayMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				x := p.Setx(count)
				for j := range x {
					x[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			case 2:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				y := p.Sety(count)
				for j := range y {
					y[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			default:
				if schema == nil || !schema.SkipCompoundArrayMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *CompoundArrayMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Field("x")
		printer.StartArray()
		for _, _it := range p._data_x{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Gety() != nil {
		printer.Field("y")
		printer.StartArray()
		for _, _it := range p._data_y{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

func (p *RecursiveMessage) Getx() *RecursiveMessage {
	if p == nil {
		return nil
	}
	return p._data_x
}

func (p *RecursiveMessage) Setx(value *RecursiveMessage) {
	p._data_x = value;
}

func (p *RecursiveMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Getx() != nil {
		bb.WriteVarUint(1)
		if !p._data_x.Encode(bb) {
			return false
		}
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *RecursiveMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_x, ok = func() (val *RecursiveMessage, ok bool) {
					v := &RecursiveMessage{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			default:
				if schema == nil || !schema.SkipRecursiveMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *RecursiveMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Getx() != nil {
		printer.Print("x", p.Getx())
	}
	printer.EndObject()
	return true;
}

func (p *NonDeprecatedMessage) Geta() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_a
	}
	return nil
}

func (p *NonDeprecatedMessage) Seta(value uint32) {
	p._flags[0] |= 1
	p._data_a = value
}

func (p *NonDeprecatedMessage) Getb() *uint32 {
	if p._flags[0] & 2 > 0 {
		return &p._data_b
	}
	return nil
}

func (p *NonDeprecatedMessage) Setb(value uint32) {
	p._flags[0] |= 2
	p._data_b = value
}

func (p *NonDeprecatedMessage) Getc() []uint32 {
	if p._flags[0] & 4 > 0 {
		return p._data_c
	}
	return nil
}

func (p *NonDeprecatedMessage) Setc(count uint32) []uint32 {
	p._flags[0] |= 4
	p._data_c = make([]uint32, int(count))
	return p._data_c
}

func (p *NonDeprecatedMessage) Getd() []uint32 {
	if p._flags[0] & 8 > 0 {
		return p._data_d
	}
	return nil
}

func (p *NonDeprecatedMessage) Setd(count uint32) []uint32 {
	p._flags[0] |= 8
	p._data_d = make([]uint32, int(count))
	return p._data_d
}

func (p *NonDeprecatedMessage) Gete() *ByteStruct {
	if p == nil {
		return nil
	}
	return p._data_e
}

func (p *NonDeprecatedMessage) Sete(value *ByteStruct) {
	p._data_e = value;
}

func (p *NonDeprecatedMessage) Getf() *ByteStruct {
	if p == nil {
		return nil
	}
	return p._data_f
}

func (p *NonDeprecatedMessage) Setf(value *ByteStruct) {
	p._data_f = value;
}

func (p *NonDeprecatedMessage) Getg() *uint32 {
	if p._flags[0] & 64 > 0 {
		return &p._data_g
	}
	return nil
}

func (p *NonDeprecatedMessage) Setg(value uint32) {
	p._flags[0] |= 64
	p._data_g = value
}

func (p *NonDeprecatedMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Geta() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(p._data_a)
	}
	if p.Getb() != nil {
		bb.WriteVarUint(2)
		bb.WriteVarUint(p._data_b)
	}
	if p.Getc() != nil {
		bb.WriteVarUint(3)
		bb.WriteVarUint(uint32(len(p._data_c)))
		for _, _it := range p._data_c{
			bb.WriteVarUint(_it)
		}
	}
	if p.Getd() != nil {
		bb.WriteVarUint(4)
		bb.WriteVarUint(uint32(len(p._data_d)))
		for _, _it := range p._data_d{
			bb.WriteVarUint(_it)
		}
	}
	if p.Gete() != nil {
		bb.WriteVarUint(5)
		if !p._data_e.Encode(bb) {
			return false
		}
	}
	if p.Getf() != nil {
		bb.WriteVarUint(6)
		if !p._data_f.Encode(bb) {
			return false
		}
	}
	if p.Getg() != nil {
		bb.WriteVarUint(7)
		bb.WriteVarUint(p._data_g)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *NonDeprecatedMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_a, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Seta(p._data_a)
				break
			case 2:
				if p._data_b, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setb(p._data_b)
				break
			case 3:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				c := p.Setc(count)
				for j := range c {
					c[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			case 4:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				d := p.Setd(count)
				for j := range d {
					d[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			case 5:
				if p._data_e, ok = func() (val *ByteStruct, ok bool) {
					v := &ByteStruct{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			case 6:
				if p._data_f, ok = func() (val *ByteStruct, ok bool) {
					v := &ByteStruct{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			case 7:
				if p._data_g, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setg(p._data_g)
				break
			default:
				if schema == nil || !schema.SkipNonDeprecatedMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *NonDeprecatedMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Geta() != nil {
		printer.Print("a", p.Geta())
	}
	if p.Getb() != nil {
		printer.Print("b", p.Getb())
	}
	if p.Getc() != nil {
		printer.Field("c")
		printer.StartArray()
		for _, _it := range p._data_c{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Getd() != nil {
		printer.Field("d")
		printer.StartArray()
		for _, _it := range p._data_d{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Gete() != nil {
		printer.Print("e", p.Gete())
	}
	if p.Getf() != nil {
		printer.Print("f", p.Getf())
	}
	if p.Getg() != nil {
		printer.Print("g", p.Getg())
	}
	printer.EndObject()
	return true;
}

func (p *DeprecatedMessage) Geta() *uint32 {
	if p._flags[0] & 1 > 0 {
		return &p._data_a
	}
	return nil
}

func (p *DeprecatedMessage) Seta(value uint32) {
	p._flags[0] |= 1
	p._data_a = value
}

func (p *DeprecatedMessage) Getc() []uint32 {
	if p._flags[0] & 4 > 0 {
		return p._data_c
	}
	return nil
}

func (p *DeprecatedMessage) Setc(count uint32) []uint32 {
	p._flags[0] |= 4
	p._data_c = make([]uint32, int(count))
	return p._data_c
}

func (p *DeprecatedMessage) Gete() *ByteStruct {
	if p == nil {
		return nil
	}
	return p._data_e
}

func (p *DeprecatedMessage) Sete(value *ByteStruct) {
	p._data_e = value;
}

func (p *DeprecatedMessage) Getg() *uint32 {
	if p._flags[0] & 64 > 0 {
		return &p._data_g
	}
	return nil
}

func (p *DeprecatedMessage) Setg(value uint32) {
	p._flags[0] |= 64
	p._data_g = value
}

func (p *DeprecatedMessage) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Geta() != nil {
		bb.WriteVarUint(1)
		bb.WriteVarUint(p._data_a)
	}
	if p.Getc() != nil {
		bb.WriteVarUint(3)
		bb.WriteVarUint(uint32(len(p._data_c)))
		for _, _it := range p._data_c{
			bb.WriteVarUint(_it)
		}
	}
	if p.Gete() != nil {
		bb.WriteVarUint(5)
		if !p._data_e.Encode(bb) {
			return false
		}
	}
	if p.Getg() != nil {
		bb.WriteVarUint(7)
		bb.WriteVarUint(p._data_g)
	}
	bb.WriteVarUint(0)
	return true;
}

func (p *DeprecatedMessage) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	highest := uint32(0)
	for {
		typ, ok := bb.ReadVarUint()
		if !ok {
			return false
		}
		if typ < highest && typ > 0 {
			return false
		}
		highest = typ
		switch typ {
			case 0:
				return true
			case 1:
				if p._data_a, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Seta(p._data_a)
				break
			case 2:
				if _, ok = bb.ReadVarUint(); !ok {
					return false
				}
				break
			case 3:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				c := p.Setc(count)
				for j := range c {
					c[j], ok = bb.ReadVarUint()
					if !ok {
						return false
					}
				}
				break
			case 4:
				count, ok = bb.ReadVarUint()
				if !ok {
					return false
				}
				for j := uint32(0); j < count; j++ {
					if _, ok = bb.ReadVarUint(); !ok {
						return false
					}
				}
				break
			case 5:
				if p._data_e, ok = func() (val *ByteStruct, ok bool) {
					v := &ByteStruct{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			case 6:
				if _, ok = func() (val *ByteStruct, ok bool) {
					v := &ByteStruct{}
					if ok = v.Decode(bb, schema); ok {
						val = v
					}
					return
				}(); !ok {
					return false
				}
				break
			case 7:
				if p._data_g, ok = bb.ReadVarUint(); !ok {
					return false
				}
				p.Setg(p._data_g)
				break
			default:
				if schema == nil || !schema.SkipDeprecatedMessageField(bb, typ) {
					return false
				}
				break
		}
	}
}

func (p *DeprecatedMessage) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Geta() != nil {
		printer.Print("a", p.Geta())
	}
	if p.Getc() != nil {
		printer.Field("c")
		printer.StartArray()
		for _, _it := range p._data_c{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Gete() != nil {
		printer.Print("e", p.Gete())
	}
	if p.Getg() != nil {
		printer.Print("g", p.Getg())
	}
	printer.EndObject()
	return true;
}

func (p *SortedStruct) Geta1() *bool {
	if p._flags[0] & 1 > 0 {
		return &p._data_a1
	}
	return nil
}

func (p *SortedStruct) Seta1(value bool) {
	p._flags[0] |= 1
	p._data_a1 = value
}

func (p *SortedStruct) Getb1() *byte {
	if p._flags[0] & 2 > 0 {
		return &p._data_b1
	}
	return nil
}

func (p *SortedStruct) Setb1(value byte) {
	p._flags[0] |= 2
	p._data_b1 = value
}

func (p *SortedStruct) Getc1() *int32 {
	if p._flags[0] & 4 > 0 {
		return &p._data_c1
	}
	return nil
}

func (p *SortedStruct) Setc1(value int32) {
	p._flags[0] |= 4
	p._data_c1 = value
}

func (p *SortedStruct) Getd1() *uint32 {
	if p._flags[0] & 8 > 0 {
		return &p._data_d1
	}
	return nil
}

func (p *SortedStruct) Setd1(value uint32) {
	p._flags[0] |= 8
	p._data_d1 = value
}

func (p *SortedStruct) Gete1() *float32 {
	if p._flags[0] & 16 > 0 {
		return &p._data_e1
	}
	return nil
}

func (p *SortedStruct) Sete1(value float32) {
	p._flags[0] |= 16
	p._data_e1 = value
}

func (p *SortedStruct) Getf1() *string {
	if p._flags[0] & 32 > 0 {
		return &p._data_f1
	}
	return nil
}

func (p *SortedStruct) Setf1(value string) {
	p._flags[0] |= 32
	p._data_f1 = value
}

func (p *SortedStruct) Geta2() *bool {
	if p._flags[0] & 64 > 0 {
		return &p._data_a2
	}
	return nil
}

func (p *SortedStruct) Seta2(value bool) {
	p._flags[0] |= 64
	p._data_a2 = value
}

func (p *SortedStruct) Getb2() *byte {
	if p._flags[0] & 128 > 0 {
		return &p._data_b2
	}
	return nil
}

func (p *SortedStruct) Setb2(value byte) {
	p._flags[0] |= 128
	p._data_b2 = value
}

func (p *SortedStruct) Getc2() *int32 {
	if p._flags[0] & 256 > 0 {
		return &p._data_c2
	}
	return nil
}

func (p *SortedStruct) Setc2(value int32) {
	p._flags[0] |= 256
	p._data_c2 = value
}

func (p *SortedStruct) Getd2() *uint32 {
	if p._flags[0] & 512 > 0 {
		return &p._data_d2
	}
	return nil
}

func (p *SortedStruct) Setd2(value uint32) {
	p._flags[0] |= 512
	p._data_d2 = value
}

func (p *SortedStruct) Gete2() *float32 {
	if p._flags[0] & 1024 > 0 {
		return &p._data_e2
	}
	return nil
}

func (p *SortedStruct) Sete2(value float32) {
	p._flags[0] |= 1024
	p._data_e2 = value
}

func (p *SortedStruct) Getf2() *string {
	if p._flags[0] & 2048 > 0 {
		return &p._data_f2
	}
	return nil
}

func (p *SortedStruct) Setf2(value string) {
	p._flags[0] |= 2048
	p._data_f2 = value
}

func (p *SortedStruct) Geta3() []bool {
	if p._flags[0] & 4096 > 0 {
		return p._data_a3
	}
	return nil
}

func (p *SortedStruct) Seta3(count uint32) []bool {
	p._flags[0] |= 4096
	p._data_a3 = make([]bool, int(count))
	return p._data_a3
}

func (p *SortedStruct) Getb3() []byte {
	if p._flags[0] & 8192 > 0 {
		return p._data_b3
	}
	return nil
}

func (p *SortedStruct) Setb3(count uint32) []byte {
	p._flags[0] |= 8192
	p._data_b3 = make([]byte, int(count))
	return p._data_b3
}

func (p *SortedStruct) Getc3() []int32 {
	if p._flags[0] & 16384 > 0 {
		return p._data_c3
	}
	return nil
}

func (p *SortedStruct) Setc3(count uint32) []int32 {
	p._flags[0] |= 16384
	p._data_c3 = make([]int32, int(count))
	return p._data_c3
}

func (p *SortedStruct) Getd3() []uint32 {
	if p._flags[0] & 32768 > 0 {
		return p._data_d3
	}
	return nil
}

func (p *SortedStruct) Setd3(count uint32) []uint32 {
	p._flags[0] |= 32768
	p._data_d3 = make([]uint32, int(count))
	return p._data_d3
}

func (p *SortedStruct) Gete3() []float32 {
	if p._flags[0] & 65536 > 0 {
		return p._data_e3
	}
	return nil
}

func (p *SortedStruct) Sete3(count uint32) []float32 {
	p._flags[0] |= 65536
	p._data_e3 = make([]float32, int(count))
	return p._data_e3
}

func (p *SortedStruct) Getf3() []string {
	if p._flags[0] & 131072 > 0 {
		return p._data_f3
	}
	return nil
}

func (p *SortedStruct) Setf3(count uint32) []string {
	p._flags[0] |= 131072
	p._data_f3 = make([]string, int(count))
	return p._data_f3
}

func (p *SortedStruct) Encode(bb *kiwi.ByteBuffer) bool {
	if p.Geta1() == nil {
		return false
	}
	bb.WriteBool(p._data_a1)
	if p.Getb1() == nil {
		return false
	}
	bb.WriteByte(p._data_b1)
	if p.Getc1() == nil {
		return false
	}
	bb.WriteVarInt(p._data_c1)
	if p.Getd1() == nil {
		return false
	}
	bb.WriteVarUint(p._data_d1)
	if p.Gete1() == nil {
		return false
	}
	bb.WriteVarFloat(p._data_e1)
	if p.Getf1() == nil {
		return false
	}
	bb.WriteString(p._data_f1)
	if p.Geta2() == nil {
		return false
	}
	bb.WriteBool(p._data_a2)
	if p.Getb2() == nil {
		return false
	}
	bb.WriteByte(p._data_b2)
	if p.Getc2() == nil {
		return false
	}
	bb.WriteVarInt(p._data_c2)
	if p.Getd2() == nil {
		return false
	}
	bb.WriteVarUint(p._data_d2)
	if p.Gete2() == nil {
		return false
	}
	bb.WriteVarFloat(p._data_e2)
	if p.Getf2() == nil {
		return false
	}
	bb.WriteString(p._data_f2)
	if p.Geta3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_a3)))
	for _, _it := range p._data_a3{
		bb.WriteBool(_it)
	}
	if p.Getb3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_b3)))
	for _, _it := range p._data_b3{
		bb.WriteByte(_it)
	}
	if p.Getc3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_c3)))
	for _, _it := range p._data_c3{
		bb.WriteVarInt(_it)
	}
	if p.Getd3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_d3)))
	for _, _it := range p._data_d3{
		bb.WriteVarUint(_it)
	}
	if p.Gete3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_e3)))
	for _, _it := range p._data_e3{
		bb.WriteVarFloat(_it)
	}
	if p.Getf3() == nil {
		return false
	}
	bb.WriteVarUint(uint32(len(p._data_f3)))
	for _, _it := range p._data_f3{
		bb.WriteString(_it)
	}
	return true;
}

func (p *SortedStruct) Decode(bb *kiwi.ByteBuffer, schema *BinarySchema) bool {
	var count uint32
	var ok bool
	_ = ok
	if p._data_a1, ok = bb.ReadBool(); !ok {
		return false
	}
	p.Seta1(p._data_a1)
	if p._data_b1, ok = bb.ReadByte(); !ok {
		return false
	}
	p.Setb1(p._data_b1)
	if p._data_c1, ok = bb.ReadVarInt(); !ok {
		return false
	}
	p.Setc1(p._data_c1)
	if p._data_d1, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Setd1(p._data_d1)
	if p._data_e1, ok = bb.ReadVarFloat(); !ok {
		return false
	}
	p.Sete1(p._data_e1)
	if p._data_f1, ok = bb.ReadString(); !ok {
		return false
	}
	p.Setf1(p._data_f1)
	if p._data_a2, ok = bb.ReadBool(); !ok {
		return false
	}
	p.Seta2(p._data_a2)
	if p._data_b2, ok = bb.ReadByte(); !ok {
		return false
	}
	p.Setb2(p._data_b2)
	if p._data_c2, ok = bb.ReadVarInt(); !ok {
		return false
	}
	p.Setc2(p._data_c2)
	if p._data_d2, ok = bb.ReadVarUint(); !ok {
		return false
	}
	p.Setd2(p._data_d2)
	if p._data_e2, ok = bb.ReadVarFloat(); !ok {
		return false
	}
	p.Sete2(p._data_e2)
	if p._data_f2, ok = bb.ReadString(); !ok {
		return false
	}
	p.Setf2(p._data_f2)
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	a3 := p.Seta3(count)
	for j := range a3 {
		a3[j], ok = bb.ReadBool()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	b3 := p.Setb3(count)
	for j := range b3 {
		b3[j], ok = bb.ReadByte()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	c3 := p.Setc3(count)
	for j := range c3 {
		c3[j], ok = bb.ReadVarInt()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	d3 := p.Setd3(count)
	for j := range d3 {
		d3[j], ok = bb.ReadVarUint()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	e3 := p.Sete3(count)
	for j := range e3 {
		e3[j], ok = bb.ReadVarFloat()
		if !ok {
			return false
		}
	}
	count, ok = bb.ReadVarUint()
	if !ok {
		return false
	}
	f3 := p.Setf3(count)
	for j := range f3 {
		f3[j], ok = bb.ReadString()
		if !ok {
			return false
		}
	}
	return true
}

func (p *SortedStruct) Print(printer *kiwi.Printer) bool {
	if p == nil { return true }
	printer = printer.With(p)
	printer.StartObject()
	if p.Geta1() != nil {
		printer.Print("a1", p.Geta1())
	}
	if p.Getb1() != nil {
		printer.Print("b1", p.Getb1())
	}
	if p.Getc1() != nil {
		printer.Print("c1", p.Getc1())
	}
	if p.Getd1() != nil {
		printer.Print("d1", p.Getd1())
	}
	if p.Gete1() != nil {
		printer.Print("e1", p.Gete1())
	}
	if p.Getf1() != nil {
		printer.Print("f1", p.Getf1())
	}
	if p.Geta2() != nil {
		printer.Print("a2", p.Geta2())
	}
	if p.Getb2() != nil {
		printer.Print("b2", p.Getb2())
	}
	if p.Getc2() != nil {
		printer.Print("c2", p.Getc2())
	}
	if p.Getd2() != nil {
		printer.Print("d2", p.Getd2())
	}
	if p.Gete2() != nil {
		printer.Print("e2", p.Gete2())
	}
	if p.Getf2() != nil {
		printer.Print("f2", p.Getf2())
	}
	if p.Geta3() != nil {
		printer.Field("a3")
		printer.StartArray()
		for _, _it := range p._data_a3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Getb3() != nil {
		printer.Field("b3")
		printer.StartArray()
		for _, _it := range p._data_b3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Getc3() != nil {
		printer.Field("c3")
		printer.StartArray()
		for _, _it := range p._data_c3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Getd3() != nil {
		printer.Field("d3")
		printer.StartArray()
		for _, _it := range p._data_d3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Gete3() != nil {
		printer.Field("e3")
		printer.StartArray()
		for _, _it := range p._data_e3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	if p.Getf3() != nil {
		printer.Field("f3")
		printer.StartArray()
		for _, _it := range p._data_f3{
			printer.Print("", &_it)
		}
		printer.EndArray()
	}
	printer.EndObject()
	return true;
}

