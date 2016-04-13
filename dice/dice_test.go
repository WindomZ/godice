package dice

import "testing"

func TestDiceIntThrow(t *testing.T) {
	d := NewDiceInt(50, TYPE_POLL)
	for i := 0; i < 100; i++ {
		t.Log(d.TV())
	}
}
