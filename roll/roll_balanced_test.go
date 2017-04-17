package roll

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func Test_RollBalanced_Roll(t *testing.T) {
	roll := NewBalancedRoll()

	assert.Empty(t, roll.Roll())

	assert.NotEmpty(t,
		roll.AddRoll(int8(1), 1).AddRoll(int16(2), 2).AddRoll(int32(3), 3).AddRoll(int64(4), 4))
	assert.NotEmpty(t,
		roll.AddRoll("a", 1).AddRoll("bb", 2).AddRoll("ccc", 3))

	check := make(map[interface{}]bool)
	for i := 0; i < roll.Size()*2; i++ {
		inst := roll.Roll()
		if _, ok := check[inst]; !ok {
			check[inst] = true
		}
	}
	assert.Equal(t, roll.Size(), len(check))

	check = make(map[interface{}]bool)
	for i := 0; i < roll.Size()*10; i++ {
		inst := roll.Roll()
		if _, ok := check[inst]; !ok {
			check[inst] = true
		}
	}
	assert.Equal(t, roll.Size(), len(check))
}
