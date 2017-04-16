package roll

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func Test_Roll_StringRoll(t *testing.T) {
	roll := NewStringRoll()

	assert.NotEmpty(t,
		roll.AddRoll(int8(1), 1).AddRoll(int16(2), 2).AddRoll(int32(3), 3).AddRoll(int64(4), 4))
	assert.NotEmpty(t,
		roll.AddRoll("a", 1).AddRoll("bb", 2).AddRoll("ccc", 3))

	for i := 0; i < 1000; i++ {
		assert.NotZero(t, len(roll.MustStringRoll()))
	}
}
