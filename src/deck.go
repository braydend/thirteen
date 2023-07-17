package thirteen

import (
	"math/rand"
	"sort"
)

type Deck = []Card

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
