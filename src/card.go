package card

type Suit = uint8
type Value = uint8

const (
	THREE = 3
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
	TWO
)

const (
	SPADE = iota
	CLUB
	DIAMOND
	HEART
)

type Card struct {
	Suit  Suit
	Value Value
}

func NewCard(value Value, suit Suit) Card {
	return Card{Suit: suit, Value: value}
}
