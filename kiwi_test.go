package kiwi_test

import (
	"testing"

	"math"

	"bytes"

	"github.com/stretchr/testify/assert"
	"github.com/petercgrant/kiwi"
	"github.com/petercgrant/kiwi/test"
)

var schema = &test.BinarySchema{}

func aTestBoolStruct(t *testing.T) {
	t.Parallel()
	check := func(i bool, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.BoolStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.BoolStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(false, []byte{0})
	check(true, []byte{1})
}

func aTestByteStruct(t *testing.T) {
	t.Parallel()
	check := func(i byte, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.ByteStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.ByteStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(0x00, []byte{0x00})
	check(0x01, []byte{0x01})
	check(0x7F, []byte{0x7F})
	check(0x80, []byte{0x80})
	check(0xFF, []byte{0xFF})
}

func aTestUintStruct(t *testing.T) {
	t.Parallel()
	check := func(i uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.UintStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.UintStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(0x00, []byte{0x00})
	check(0x01, []byte{0x01})
	check(0x02, []byte{0x02})
	check(0x7F, []byte{0x7F})
	check(0x80, []byte{0x80, 0x01})
	check(0x81, []byte{0x81, 0x01})
	check(0x100, []byte{0x80, 0x02})
	check(0x101, []byte{0x81, 0x02})
	check(0x17F, []byte{0xFF, 0x02})
	check(0x180, []byte{0x80, 0x03})
	check(0x1FF, []byte{0xFF, 0x03})
	check(0x200, []byte{0x80, 0x04})
	check(0x7FFF, []byte{0xFF, 0xFF, 0x01})
	check(0x8000, []byte{0x80, 0x80, 0x02})
	check(0x7FFFFFFF, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x07})
	check(0x80000000, []byte{0x80, 0x80, 0x80, 0x80, 0x08})
}

func aTestIntStruct(t *testing.T) {
	t.Parallel()
	check := func(i int32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.IntStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.IntStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(0x00, []byte{0x00})
	check(-0x01, []byte{0x01})
	check(0x01, []byte{0x02})
	check(-0x02, []byte{0x03})
	check(0x02, []byte{0x04})
	check(-0x3F, []byte{0x7D})
	check(0x3F, []byte{0x7E})
	check(-0x40, []byte{0x7F})
	check(0x40, []byte{0x80, 0x01})
	check(-0x3FFF, []byte{0xFD, 0xFF, 0x01})
	check(0x3FFF, []byte{0xFE, 0xFF, 0x01})
	check(-0x4000, []byte{0xFF, 0xFF, 0x01})
	check(0x4000, []byte{0x80, 0x80, 0x02})
	check(-0x3FFFFFFF, []byte{0xFD, 0xFF, 0xFF, 0xFF, 0x07})
	check(0x3FFFFFFF, []byte{0xFE, 0xFF, 0xFF, 0xFF, 0x07})
	check(-0x40000000, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x07})
	check(0x40000000, []byte{0x80, 0x80, 0x80, 0x80, 0x08})
	check(-0x7FFFFFFF, []byte{0xFD, 0xFF, 0xFF, 0xFF, 0x0F})
	check(0x7FFFFFFF, []byte{0xFE, 0xFF, 0xFF, 0xFF, 0x0F})
	check(-0x80000000, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x0F})
}

func aTestFloatStruct(t *testing.T) {
	t.Parallel()
	check := func(i float32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.FloatStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.FloatStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		if x := z.Getx(); x != nil && math.IsNaN(float64(*x)) {
			assert.Equal(t, z.Getx() == nil, v.Getx() == nil)
			assert.Equal(t, math.IsNaN(float64(*z.Getx())), math.IsNaN(float64(*v.Getx())))
		} else {
			assert.EqualValues(t, z, v)
		}
	}
	check(0, []byte{0})
	check(1, []byte{127, 0, 0, 0})
	check(-1, []byte{127, 1, 0, 0})
	check(3.1415927410125732, []byte{128, 182, 31, 146})
	check(-3.1415927410125732, []byte{128, 183, 31, 146})
	check(float32(math.Inf(1)), []byte{255, 0, 0, 0})
	check(float32(math.Inf(-1)), []byte{255, 1, 0, 0})
	check(float32(math.NaN()), []byte{255, 0, 0, 128})
}

func aTestStringStruct(t *testing.T) {
	t.Parallel()
	check := func(i string, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.StringStruct{}
		v.Setx(i)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.StringStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check("", []byte{0})
	check("abc", []byte{97, 98, 99, 0})
	check("🙉🙈🙊", []byte{240, 159, 153, 137, 240, 159, 153, 136, 240, 159, 153, 138, 0})
}

func aTestCompoundStruct(t *testing.T) {
	t.Parallel()
	check := func(x, y uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.CompoundStruct{}
		v.Setx(x)
		v.Sety(y)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.CompoundStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(0, 0, []byte{0, 0})
	check(1, 2, []byte{1, 2})
	check(12345, 54321, []byte{185, 96, 177, 168, 3})
}

func aTestNestedStruct(t *testing.T) {
	t.Parallel()
	check := func(a, x, y, c uint32, d map[string][]uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.NestedStruct{}
		v.Seta(a)
		cs := &test.CompoundStruct{}
		cs.Setx(x)
		cs.Sety(y)
		v.Setb(cs)
		v.Setc(c)
		mm := make(map[string]*test.CompoundStruct, len(d))
		for k, v := range d {
			m := &test.CompoundStruct{}
			m.Setx(v[0])
			m.Sety(v[1])
			mm[k] = m
		}
		v.Setd(mm)
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.NestedStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(0, 0, 0, 0, map[string][]uint32{"abc": []uint32{1, 2}}, []byte{0, 0, 0, 0, 1, 97, 98, 99, 0, 1, 2})
	check(1, 2, 3, 4, map[string][]uint32{}, []byte{1, 2, 3, 4, 0})
	check(534, 12345, 54321, 321, map[string][]uint32{}, []byte{150, 4, 185, 96, 177, 168, 3, 193, 2, 0})
}

func aTestEnumMessage(t *testing.T) {
	t.Parallel()
	check := func(i *test.Enum, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.EnumMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.EnumMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	a := test.A
	check(nil, []byte{0})
	check(&a, []byte{1, 100, 0})

	z := &test.EnumMessage{}
	assert.False(t, z.Decode(kiwi.NewSharedByteBuffer([]byte{1, 244, 0}), schema))
}

func aTestBoolMessage(t *testing.T) {
	t.Parallel()
	check := func(i *bool, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.BoolMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.BoolMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	pf := false
	pt := true
	check(&pf, []byte{1, 0, 0})
	check(&pt, []byte{1, 1, 0})
}

func aTestByteMessage(t *testing.T) {
	t.Parallel()
	check := func(i *byte, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.ByteMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.ByteMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	pb := byte(234)
	check(&pb, []byte{1, 234, 0})
}

func aTestUintMessage(t *testing.T) {
	t.Parallel()
	check := func(i *uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.UintMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.UintMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	px := uint32(12345678)
	check(&px, []byte{1, 206, 194, 241, 5, 0})
}

func aTestIntMessage(t *testing.T) {
	t.Parallel()
	check := func(i *int32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.IntMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.IntMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	px := int32(12345678)
	check(&px, []byte{1, 156, 133, 227, 11, 0})
}

func aTestFloatMessage(t *testing.T) {
	t.Parallel()
	check := func(i *float32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.FloatMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.FloatMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	px := float32(3.1415927410125732)
	check(&px, []byte{1, 128, 182, 31, 146, 0})
}

func aTestStringMessage(t *testing.T) {
	t.Parallel()
	check := func(i *string, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.StringMessage{}
		if i != nil {
			v.Setx(*i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.StringMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	empty := ""
	check(&empty, []byte{1, 0, 0})
	px := "🙉🙈🙊"
	check(&px, []byte{1, 240, 159, 153, 137, 240, 159, 153, 136, 240, 159, 153, 138, 0, 0})
}

func aTestMapMessage(t *testing.T) {
	t.Parallel()
	check := func(i map[string]int32, oo ...[]byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.MapMessage{}
		if i != nil {
			v.Setx(i)
		}
		assert.True(t, v.Encode(enc))
		found := false
		bb := enc.Bytes()
		for _, o := range oo {
			if bytes.Equal(o, bb) {
				found = true
				break
			}
		}
		assert.True(t, found)

		z := &test.MapMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(bb), schema))
		mm := z.Getx()
		nn := v.Getx()
		if mm == nil || nn == nil {
			assert.Equal(t, z, v)
		} else {
			for a, b := range *mm {
				assert.EqualValues(t, b, (*nn)[a])
			}
		}
	}
	check(nil, []byte{0})
	// go maps are not sorted
	check(map[string]int32{"abc": 5, "def": 10}, []byte{1, 2, 97, 98, 99, 0, 10, 100, 101, 102, 0, 20, 0}, []byte{1, 2, 100, 101, 102, 0, 20, 97, 98, 99, 0, 10, 0})
}

func aTestCompoundMessage(t *testing.T) {
	t.Parallel()
	check := func(x, y *uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.CompoundMessage{}
		if x != nil {
			v.Setx(*x)
		}
		if y != nil {
			v.Sety(*y)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.CompoundMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, nil, []byte{0})
	px := uint32(123)
	check(&px, nil, []byte{1, 123, 0})
	py := uint32(234)
	check(nil, &py, []byte{2, 234, 1, 0})
	check(&px, &py, []byte{1, 123, 2, 234, 1, 0})
}

func Uint32(v uint32) *uint32 {
	return &v
}

func aTestNestedMessage(t *testing.T) {
	t.Parallel()
	check := func(a, x, y, c *uint32, d map[string][]*uint32, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.NestedMessage{}
		if a != nil {
			v.Seta(*a)
		}
		cs := &test.CompoundMessage{}
		if x != nil {
			cs.Setx(*x)
		}
		if y != nil {
			cs.Sety(*y)
		}
		if x != nil || y != nil {
			v.Setb(cs)
		}
		if c != nil {
			v.Setc(*c)
		}
		mm := make(map[string]*test.CompoundMessage, len(d))
		for k, v := range d {
			m := &test.CompoundMessage{}
			if v[0] != nil {
				m.Setx(*v[0])
			}
			if v[1] != nil {
				m.Sety(*v[1])
			}
			mm[k] = m
		}
		if len(mm) > 0 {
			v.Setd(mm)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.NestedMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, nil, nil, nil, nil, []byte{0})
	check(Uint32(123), nil, nil, Uint32(234), nil, []byte{1, 123, 3, 234, 1, 0})
	check(nil, Uint32(5), Uint32(6), nil, nil, []byte{2, 1, 5, 2, 6, 0, 0})
	check(nil, Uint32(5), nil, Uint32(123), nil, []byte{2, 1, 5, 0, 3, 123, 0})
	check(Uint32(234), Uint32(5), Uint32(6), Uint32(123), map[string][]*uint32{"a": []*uint32{Uint32(1), Uint32(2)}}, []byte{1, 234, 1, 2, 1, 5, 2, 6, 0, 3, 123, 4, 1, 97, 0, 1, 1, 2, 2, 0, 0})
}

func aTestEnumArrayStruct(t *testing.T) {
	t.Parallel()
	check := func(i []test.Enum, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.EnumArrayStruct{}
		if i != nil {
			v.Setx(uint32(len(i)))
			for j := range i {
				v.Getx()[j] = i[j]
			}
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.EnumArrayStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	//check(nil, []byte{0}) // TODO: FIXME
	check([]test.Enum{}, []byte{0})
	check([]test.Enum{test.B, test.A}, []byte{2, 200, 1, 100})
}

func aTestEnumArrayMessage(t *testing.T) {
	t.Parallel()
	check := func(i []test.Enum, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.EnumArrayMessage{}
		if i != nil {
			v.Setx(uint32(len(i)))
			for j := range i {
				v.Getx()[j] = i[j]
			}
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.EnumArrayMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	check([]test.Enum{}, []byte{1, 0, 0})
	check([]test.Enum{test.B, test.A}, []byte{1, 2, 200, 1, 100, 0})
}

func aTestBoolArrayStruct(t *testing.T) {
	t.Parallel()
	check := func(i []bool, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.BoolArrayStruct{}
		if i != nil {
			v.Setx(uint32(len(i)))
			for j := range i {
				v.Getx()[j] = i[j]
			}
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.BoolArrayStruct{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	//check(nil, []byte{0}) // TODO: FIXME
	check([]bool{}, []byte{0})
	check([]bool{true, false}, []byte{2, 1, 0})
}

func aTestBoolArrayMessage(t *testing.T) {
	t.Parallel()
	check := func(i []bool, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.BoolArrayMessage{}
		if i != nil {
			v.Setx(uint32(len(i)))
			for j := range i {
				v.Getx()[j] = i[j]
			}
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.BoolArrayMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	check([]bool{}, []byte{1, 0, 0})
	check([]bool{true, false}, []byte{1, 2, 1, 0, 0})
}

func aTestRecursiveMessage(t *testing.T) {
	t.Parallel()
	check := func(i *test.RecursiveMessage, o []byte) {
		enc := kiwi.NewByteBuffer()
		v := &test.RecursiveMessage{}
		if i != nil {
			v.Setx(i)
		}
		assert.True(t, v.Encode(enc))
		assert.EqualValues(t, o, enc.Bytes())

		z := &test.RecursiveMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(o), schema))
		assert.EqualValues(t, z, v)
	}
	check(nil, []byte{0})
	px := &test.RecursiveMessage{}
	check(px, []byte{1, 0, 0})
	px.Setx(&test.RecursiveMessage{})
	check(px, []byte{1, 1, 0, 0, 0})
}

// TODO: binary schema
//it('binary schema', function () {
//var compiledSchema = kiwi.compileSchema(kiwi.encodeBinarySchema(schemaText));
//
//function check(message) {
//assert.deepEqual(
//Buffer(schema.encodeNestedMessage(message)),
//Buffer(compiledSchema.encodeNestedMessage(message)));
//}
//
//check({ a: 1, c: 4 });
//check({ a: 1, b: {}, c: 4 });
//check({ a: 1, b: { x: 2, y: 3 }, c: 4 });
//});

// TODO: large schema
//var largeSchemaText = fs.readFileSync(__dirname + '/test-schema-large.kiwi', 'utf8');
//var largeSchema = kiwi.compileSchema(largeSchemaText);
//
//it('struct with many fields', function () {
//var object = {};
//for (var i = 0; i < 130; i++) object['f' + i] = i;
//
//var encoded = largeSchema.encodeStruct(object);
//assert.deepEqual(object, largeSchema.decodeStruct(encoded));
//});
//
//it('message with many fields', function () {
//var object = {};
//for (var i = 0; i < 130; i++) object['f' + i] = i;
//
//var encoded = largeSchema.encodeMessage(object);
//assert.deepEqual(object, largeSchema.decodeMessage(encoded));
//});

func TestDeprecatedMessage(t *testing.T) {
	t.Parallel()
	check := func(v *test.DeprecatedMessage, w *test.NonDeprecatedMessage) {
		wb := kiwi.NewByteBuffer()
		assert.True(t, w.Encode(wb))
		y := &test.DeprecatedMessage{}
		assert.True(t, y.Decode(kiwi.NewSharedByteBuffer(wb.Bytes()), schema))

		assert.EqualValues(t, y, v)

		vb := kiwi.NewByteBuffer()
		assert.True(t, v.Encode(vb))

		yb := kiwi.NewByteBuffer()
		assert.True(t, y.Encode(yb))
		//assert.EqualValues(t, yb.Bytes(), vb.Bytes())

		z := &test.NonDeprecatedMessage{}
		assert.True(t, z.Decode(kiwi.NewSharedByteBuffer(vb.Bytes()), schema))

		zb := kiwi.NewByteBuffer()
		assert.True(t, z.Encode(zb))
		assert.EqualValues(t, zb.Bytes(), vb.Bytes())
		assert.NotEqual(t, w, v)
	}
	nonDeprecated := &test.NonDeprecatedMessage{}
	nonDeprecated.Seta(1)
	nonDeprecated.Setb(2)
	c := nonDeprecated.Setc(3)
	copy(c, []uint32{3, 4, 5})
	d := nonDeprecated.Setd(3)
	copy(d, []uint32{6, 7, 8})
	be := &test.ByteStruct{}
	be.Setx(123)
	nonDeprecated.Sete(be)
	bf := &test.ByteStruct{}
	bf.Setx(13)
	nonDeprecated.Setf(bf)
	nonDeprecated.Setg(9)

	deprecated := &test.DeprecatedMessage{}
	deprecated.Seta(1)
	c = deprecated.Setc(3)
	copy(c, []uint32{3, 4, 5})
	bs := &test.ByteStruct{}
	bs.Setx(123)
	deprecated.Sete(bs)
	deprecated.Setg(9)

	check(deprecated, nonDeprecated)
}
