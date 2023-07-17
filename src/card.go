package thirteen

import (
	"fmt"
)

type Suit = uint8
type Value = uint8

const (
	THREE Suit = 3
	FOUR  Suit = 4
	FIVE  Suit = 5
	SIX   Suit = 6
	SEVEN Suit = 7
	EIGHT Suit = 8
	NINE  Suit = 9
	TEN   Suit = 10
	JACK  Suit = 11
	QUEEN Suit = 12
	KING  Suit = 13
	ACE   Suit = 14
	TWO   Suit = 15
)

const (
	SPADE Suit = iota
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

// TODO: Use stringer to generate these
func (card *Card) SuitName() string {
	switch card.Suit {
	case DIAMOND:
		return "Diamond"
	case HEART:
		return "Heart"
	case SPADE:
		return "Spade"
	case CLUB:
		return "Club"
	}

	return ""
}

func (card *Card) ValueName() string {
	switch card.Value {
	case THREE:
		return "Three"
	case FOUR:
		return "Four"
	case FIVE:
		return "Five"
	case SIX:
		return "Six"
	case SEVEN:
		return "Seven"
	case EIGHT:
		return "Eight"
	case NINE:
		return "Nine"
	case TEN:
		return "Ten"
	case JACK:
		return "Jack"
	case QUEEN:
		return "Queen"
	case KING:
		return "King"
	case ACE:
		return "Ace"
	case TWO:
		return "Two"
	}

	return ""
}

func (card *Card) ToString() string {
	suit := card.SuitName()
	value := card.ValueName()

	return fmt.Sprintf("%s of %ss", value, suit)
}
