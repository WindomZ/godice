package roll

import "fmt"

type StringRoll interface {
	Roll
	StringRoll() (string, bool)
	MustStringRoll() string
}

func NewStringRoll() StringRoll {
	return &_StringRoll{
		_Roll: *newRoll(),
	}
}

type _StringRoll struct {
	_Roll
}

func (r _Roll) StringRoll() (string, bool) {
	if inst := r.Roll(); inst == nil {
	} else if s, ok := inst.(string); ok {
		return s, true
	} else {
		return fmt.Sprint(inst), true
	}
	return "", false
}

func (r _Roll) MustStringRoll() string {
	s, _ := r.StringRoll()
	return s
}
