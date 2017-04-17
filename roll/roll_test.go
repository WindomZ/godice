package roll

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewRoll(t *testing.T) {
	roll := NewRoll()
	assert.NotEmpty(t, roll)
}

func Test_Roll_Weight(t *testing.T) {
	roll := NewRoll()

	assert.NotEmpty(t, roll.AddRoll(-1, -1).AddRoll(0, 0))
	assert.NotEmpty(t, roll.AddRoll(int8(1), 1).AddRoll(int16(2), 2).AddRoll(int32(3), 3).AddRoll(int64(4), 4))
	assert.NotEmpty(t, roll.AddRoll("a", 1).AddRoll("bb", 2).AddRoll("ccc", 3))
	assert.NotEmpty(t, roll.AddRoll(nil, 1).AddRoll(nil, 2).AddRoll(nil, 3))

	assert.Equal(t, roll.Weight(-1), 0)
	assert.Equal(t, roll.Weight(0), 0)
	assert.Equal(t, roll.Weight(int8(1)), 1)
	assert.Equal(t, roll.Weight(int16(2)), 2)
	assert.Equal(t, roll.Weight(int32(3)), 3)
	assert.Equal(t, roll.Weight(int64(4)), 4)

	assert.Equal(t, roll.Weight("a"), 1)
	assert.Equal(t, roll.Weight("bb"), 2)
	assert.Equal(t, roll.Weight("ccc"), 3)

	assert.Equal(t, roll.Weight(nil), 6)
}

func Test_Roll_Roll(t *testing.T) {
	roll := NewRoll()

	assert.Empty(t, roll.Roll())

	assert.NotEmpty(t,
		roll.AddRoll(int8(1), 1).AddRoll(int16(2), 2).AddRoll(int32(3), 3).AddRoll(int64(4), 4))
	assert.NotEmpty(t,
		roll.AddRoll("a", 1).AddRoll("bb", 2).AddRoll("ccc", 3))

	check := make(map[interface{}]bool)
	for i := 0; i < 1000; i++ {
		inst := roll.Roll()
		if _, ok := check[inst]; !ok {
			check[inst] = true
		}
	}
	assert.Equal(t, roll.Size(), len(check))
}
