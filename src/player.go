package thirteen

import (
	"log"

	"github.com/google/uuid"
)

type Player struct {
	id    string
	name  string
	isCpu bool
	cards []Card
}

func NewPlayer(name string) Player {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatalf("Failed to generate UUID for player. %s", err)
	}

	return Player{id: uuid.String(), name: name, isCpu: true, cards: []Card{}}
}

func (player *Player) AddCard(card Card) {
	player.cards = append(player.cards, card)
}

func (player *Player) HasCard(suit Suit, value Value) bool {
	for _, card := range player.cards {
		if card.Suit == suit && card.Value == value {
			return true
		}
	}

	return false
}
