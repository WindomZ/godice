package dice

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewDice(t *testing.T) {
	d := NewDice(5, TYPE_DEFAULT)
	assert.Equal(t, d.Roll(), d)
	assert.Equal(t, d.Roll().Dice(), int64(2))
	assert.Equal(t, d.Roll().Dice(), int64(4))
}
