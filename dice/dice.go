package dice

import (
	"math/rand"
	"sync"
	"time"
)

type baseDice interface {
	Throw() interface{}
}

type DiceType uint8

const (
	TYPE_POLL DiceType = iota
	TYPE_RANDOM
	TYPE_MIXED
)

type DiceInt struct {
	baseDice
	idx   int
	Value int
	Max   int
	Type  DiceType
	mutex *sync.Mutex
}

func NewDiceInt(max int, t DiceType) *DiceInt {
	if max <= 0 {
		max = 1
	}
	rand.Seed(time.Now().Unix())
	return &DiceInt{idx: 0, Value: 0, Max: max, Type: t, mutex: &sync.Mutex{}}
}

func (s *DiceInt) Throw() *DiceInt {
	s.mutex.Lock()
	defer s.mutex.Unlock()
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

func (s *DiceInt) T() *DiceInt {
	return s.Throw()
}

func (s *DiceInt) V() int {
	return s.Value
}

func (s *DiceInt) TV() int {
	return s.Throw().Value
}
