package thirteen

import (
	"log"
	"math/rand"
	"sort"
)

type Deck = []Card

func NewDeck() (deck Deck) {
	for i := 0; i < 52; i++ {
		card, err := NewCard(Value(i%13+3), Suit(i%4+1))
		if err != nil {
			log.Fatalf("Failed to create card. %v", err)
		}
		deck = append(deck, card)
	}

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

func Deal(deck Deck, players [4]Player) {
	for i, card := range deck {
		players[i%4].AddCard(&card)
	}
}
