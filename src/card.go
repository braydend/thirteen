package card

import (
	"fmt"
	"math/rand"
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
	SPADE Suit = iota
	CLUB
	DIAMOND
	HEART
)

type Card struct {
	Suit  Suit
	Value Value
}

type Deck = []Card

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

func NewDeck() (deck Deck) {
	for i := 0; i < 52; i++ {
		deck = append(deck, NewCard(Value(i%13+3), Suit(i%4)))
	}

	return deck
}

func StringifyDeck(deck Deck) (out string) {
	for _, card := range deck {
		out = out + card.ToString() + "\n"
	}

	return out
}

func SortDeck(deck Deck) Deck {
	sort.Slice(deck, func(i, j int) bool {
		if deck[i].Suit == deck[j].Suit {
			return deck[i].Value < deck[j].Value
		}

		return deck[i].Suit < deck[j].Suit
	})

	return deck
}

func ShuffleDeck(deck Deck) Deck {
	sort.Slice(deck, func(_, _ int) bool {
		randOne := rand.Intn(5)
		randTwo := rand.Intn(5)
		return randOne < randTwo
	})

	return deck
}
