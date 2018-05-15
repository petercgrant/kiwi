package kiwi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mnk.ee/kiwi"
	"mnk.ee/kiwi/test"
)

func TestEncodeDecodeMapMessage(t *testing.T) {
	t.Parallel()

	schema := &test.BinarySchema{}
	mapMessage := &test.MapMessage{}
	mapMessage.Setx(map[string]int32{"abc": 5, "def": 10})
	newb := kiwi.NewByteBuffer()
	assert.True(t, mapMessage.Encode(newb))
	assert.EqualValues(t, []byte{1, 2, 97, 98, 99, 0, 10, 100, 101, 102, 0, 20, 0}, newb.Bytes())
	mm := &test.MapMessage{}
	assert.True(t, mm.Decode(newb, schema))
	assert.EqualValues(t, mapMessage, mm)
}

func TestEncodeDecodeMapMessage_Empty(t *testing.T) {
	t.Parallel()

	schema := &test.BinarySchema{}
	mapMessage := &test.MapMessage{}
	newb := kiwi.NewByteBuffer()
	assert.True(t, mapMessage.Encode(newb))
	assert.EqualValues(t, []byte{0}, newb.Bytes())
	mm := &test.MapMessage{}
	assert.True(t, mm.Decode(newb, schema))
	assert.EqualValues(t, mapMessage, mm)
}
