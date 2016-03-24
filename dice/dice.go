package dice

import (
	"math/rand"
	"time"
)

type baseDice interface {
	Throw() interface{}
}

const (
	TYPE_POLL int = iota
	TYPE_RANDOM
	TYPE_MIXED
)

type DiceInt struct {
	baseDice
	idx   int
	Value int
	Max   int
	Type  int
}

func NewDiceInt(max, t int) *DiceInt {
	if max <= 0 {
		max = 1
	}
	rand.Seed(time.Now().Unix())
	return &DiceInt{idx: 0, Value: 0, Max: max, Type: t}
}

func (s *DiceInt) Throw() *DiceInt {
	switch s.Type {
	case TYPE_RANDOM:
		s.Value = rand.Intn(s.Max)
	case TYPE_MIXED:
		s.Value = (s.idx + rand.Intn(s.Max)) % s.Max
		s.idx++
	default:
		s.Value = s.idx % s.Max
		s.idx++
	}
	return s
}
