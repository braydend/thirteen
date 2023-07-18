package thirteen

import (
	"log"

	"github.com/google/uuid"
)

type Player struct {
	id       string
	name     string
	isCpu    bool
	cards    *[]Card
	playMove func(cards []Card) error
}

func NewPlayer(name string, playFn func(cards []Card) error) Player {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatalf("Failed to generate UUID for player. %s", err)
	}

	return Player{id: uuid.String(), name: name, isCpu: true, cards: &[]Card{}, playMove: playFn}
}

func (player *Player) AddCard(card *Card) {
	currentCards := *player.cards
	currentCards = append(currentCards, *card)

	player.cards = &currentCards
}

func (player *Player) HasCard(suit Suit, value Value) bool {
	for _, card := range *player.cards {
		if card.Suit == suit && card.Value == value {
			return true
		}
	}

	return false
}

func (player *Player) PlayMove(cards []Card) error {
	log.Printf("%s attempting to play the following cards:\n", player.name)
	for _, card := range cards {
		log.Println(card.ToString())
	}

	err := player.playMove(cards)

	if err != nil {
		return err
	}

	for _, currentCard := range cards {
		player.RemoveCard(currentCard)
	}

	return nil
}

func (player *Player) RemoveCard(card Card) error {
	var remainingCards []Card
	for _, cardInHand := range *player.cards {
		if card.Suit != cardInHand.Suit || card.Value != cardInHand.Value {
			remainingCards = append(remainingCards, cardInHand)
			log.Printf("Keeping %s\n", cardInHand.ToString())
		}
	}

	*player.cards = remainingCards

	return nil
}
