package roll

import (
	"math/rand"
)

type _RollBalanced struct {
	_Roll
	balanced map[interface{}]int64
}

func (r *_RollBalanced) Roll() interface{} {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.balanced == nil || len(r.balanced) == 0 || r.weights <= 0 {
		r.balanced = make(map[interface{}]int64, len(r.mapping))
		r.weights = 0
		for inst, weight := range r.mapping {
			r.balanced[inst] = weight
			r.weights += weight
		}
	}
	index := rand.Int63n(r.weights)
	for inst, weight := range r.balanced {
		if index -= weight; index <= 0 {
			delete(r.balanced, inst)
			r.weights -= weight
			return inst
		}
	}
	return nil
}
