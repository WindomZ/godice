package dice

import (
	"math/rand"
	"sync"
	"time"
)

type DiceInt struct {
	idx   int64
	Value int64
	Max   int64
	Type  DiceType
	mutex *sync.Mutex
}

func NewDiceInt(max int, t DiceType) *DiceInt {
	return NewDiceInt64(int64(max), t)
}

func NewDiceInt64(max int64, t DiceType) *DiceInt {
	if max <= 0 {
		max = 1
	}
	rand.Seed(time.Now().Unix())
	return &DiceInt{
		idx:   0,
		Value: 0,
		Max:   max,
		Type:  t,
		mutex: &sync.Mutex{},
	}
}

func (s *DiceInt) roll() *DiceInt {
	s.mutex.Lock()
	switch s.Type {
	case TYPE_RANDOM:
		s.Value = rand.Int63n(s.Max)
	case TYPE_MIXED:
		s.Value = (s.idx + rand.Int63n(s.Max)) % s.Max
		s.idx++
	default:
		s.Value = s.idx % s.Max
		s.idx++
	}
	s.mutex.Unlock()
	return s
}

func (s *DiceInt) Roll() Dice {
	return s.roll()
}

func (s *DiceInt) Dice() int64 {
	return s.roll().Value
}

func (s *DiceInt) DiceInt() int {
	return int(s.Dice())
}

func (s *DiceInt) DiceInt8() int8 {
	return int8(s.Dice())
}

func (s *DiceInt) DiceInt16() int16 {
	return int16(s.Dice())
}

func (s *DiceInt) DiceInt32() int32 {
	return int32(s.Dice())
}
