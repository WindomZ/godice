package dice

type DiceType uint8

const (
	TYPE_INCREMENT DiceType = iota
	TYPE_RANDOM
	TYPE_MIXED
)

const TYPE_DEFAULT = TYPE_INCREMENT

type Dice interface {
	Roll() Dice
	Dice() int64
}

func NewDice(max int64, t DiceType) Dice {
	return NewDiceInt64(max, t)
}
