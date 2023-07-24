package thirteen

import (
	"fmt"
	"log"
	"math"
	"sort"
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
	SPADE   Suit = 1
	CLUB    Suit = 2
	DIAMOND Suit = 3
	HEART   Suit = 4
)

type Card struct {
	Suit  Suit
	Value Value
}

func NewCard(value Value, suit Suit) (Card, error) {
	if value < 3 || value > 15 {
		return Card{}, fmt.Errorf("Invalid value: %d", value)
	}

	if suit == 0 || suit > 4 {
		return Card{}, fmt.Errorf("Invalid suit: %d", value)
	}

	return Card{Suit: suit, Value: value}, nil
}

// TODO: Use stringer to generate these
func (card *Card) suitName() string {
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

func (card *Card) valueName() string {
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

func (card Card) String() string {
	suit := card.suitName()
	value := card.valueName()

	return fmt.Sprintf("%s of %ss", value, suit)
}

func SortCards(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Value == cards[j].Value {
			return cards[i].Suit < cards[j].Suit
		}

		return cards[i].Value < cards[j].Value
	})

	return cards
}

func StringifyCards(cards []Card) (out string) {
	for _, card := range cards {
		out = out + card.String() + "\n"
	}

	return out
}

func (card Card) Score() int {
	score := int(math.Pow(float64(card.Value), float64(card.Suit)))

	log.Printf("%s has the score: %d", card.String(), score)
	return score
}
