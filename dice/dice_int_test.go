package dice

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewDiceInt(t *testing.T) {
	assert.NotEmpty(t, NewDiceInt(1, TYPE_RANDOM))
}

func TestNewDiceInt64(t *testing.T) {
	assert.NotEmpty(t, NewDiceInt64(1, TYPE_MIXED))
}

func TestDiceInt_Roll(t *testing.T) {
	d := NewDiceInt(3, TYPE_INCREMENT)
	assert.Equal(t, d.Roll(), d)
	assert.Equal(t, d.Roll().Dice(), int64(2))
	assert.Equal(t, d.Roll().Dice(), int64(1))
}

func TestDiceInt_Dice(t *testing.T) {
	d := NewDiceInt(50, TYPE_INCREMENT)
	for i := 0; i < 100; i++ {
		assert.Equal(t, d.Dice(), int64(i%50))
	}
}

func TestDiceInt_DiceInt(t *testing.T) {
	d := NewDiceInt(50, TYPE_INCREMENT)
	for i := 0; i < 100; i++ {
		assert.Equal(t, d.DiceInt(), i%50)
	}
}

func TestDiceInt_DiceInt8(t *testing.T) {
	d := NewDiceInt(50, TYPE_INCREMENT)
	for i := 0; i < 100; i++ {
		assert.Equal(t, d.DiceInt8(), int8(i%50))
	}
}

func TestDiceInt_DiceInt16(t *testing.T) {
	d := NewDiceInt(50, TYPE_INCREMENT)
	for i := 0; i < 100; i++ {
		assert.Equal(t, d.DiceInt16(), int16(i%50))
	}
}

func TestDiceInt_DiceInt32(t *testing.T) {
	d := NewDiceInt(50, TYPE_INCREMENT)
	for i := 0; i < 100; i++ {
		assert.Equal(t, d.DiceInt32(), int32(i%50))
	}
}
