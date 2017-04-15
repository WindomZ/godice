package roll

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Roll interface {
	AddRoll(inst interface{}, weight int) Roll
	Size() int
	Weight(inst interface{}) int
	Roll() interface{}
}

func NewRoll() Roll {
	return newRoll()
}

func NewBalancedRoll() Roll {
	return &_RollBalanced{
		_Roll:    *newRoll(),
		balanced: make(map[interface{}]int64),
	}
}

type _Roll struct {
	mapping map[interface{}]int64
	weights int64
	mutex   *sync.RWMutex
}

func newRoll() *_Roll {
	return &_Roll{
		mapping: make(map[interface{}]int64),
		weights: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (r *_Roll) AddRoll(inst interface{}, weight int) Roll {
	if weight <= 0 {
		return r
	}
	r.mutex.Lock()
	if w, ok := r.mapping[inst]; ok {
		r.mapping[inst] = int64(weight) + w
	} else {
		r.mapping[inst] = int64(weight)
	}
	atomic.AddInt64(&r.weights, int64(weight))
	r.mutex.Unlock()
	return r
}

func (r _Roll) Size() int {
	return len(r.mapping)
}

func (r _Roll) Weight(inst interface{}) int {
	if weight, ok := r.mapping[inst]; ok {
		return int(weight)
	}
	return 0
}

func (r _Roll) Roll() interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	index := rand.Int63n(r.weights)
	for inst, weight := range r.mapping {
		if index -= weight; index <= 0 {
			return inst
		}
	}
	return nil
}
