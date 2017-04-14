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
	Weight(inst interface{}) int
	Roll() interface{}
}

func NewRoll() Roll {
	return newRoll()
}

type _Roll struct {
	mutex   *sync.RWMutex
	mapping map[interface{}]int64
	weights int64
}

func newRoll() *_Roll {
	return &_Roll{
		mutex:   &sync.RWMutex{},
		mapping: make(map[interface{}]int64),
		weights: 0,
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

func (r _Roll) Weight(inst interface{}) int {
	if w, ok := r.mapping[inst]; ok {
		return int(w)
	}
	return 0
}

func (r _Roll) Roll() interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	index := rand.Int63n(r.weights)
	for v, w := range r.mapping {
		if index -= w; index <= 0 {
			return v
		}
	}
	return nil
}
