package thirteen

import (
	"log"

	"github.com/google/uuid"
)

type UserPlayer struct {
	id       string
	name     string
	isCpu    bool
	cards    *[]Card
	playMove func(cards []Card) error
}

func NewUserPlayer(name string, playFn func(cards []Card) error) UserPlayer {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatalf("Failed to generate UUID for player. %s", err)
	}

	return UserPlayer{id: uuid.String(), name: name, isCpu: true, cards: &[]Card{}, playMove: playFn}
}

func (player UserPlayer) Id() string {
	return player.id
}

func (player UserPlayer) Name() string {
	return player.name
}

func (player UserPlayer) Cards() *[]Card {
	return player.cards
}

func (player UserPlayer) AddCard(card *Card) {
	currentCards := *player.cards
	currentCards = append(currentCards, *card)

	*player.cards = currentCards
}

func (player UserPlayer) CardCount() uint8 {
	return uint8(len(*player.cards))
}

func (player UserPlayer) Play() (bool, error) {
	// TODO: Implement this
	return true, nil
}

func (player UserPlayer) HasCard(suit Suit, value Value) bool {
	for _, card := range *player.cards {
		if card.Suit == suit && card.Value == value {
			return true
		}
	}

	return false
}

func (player *UserPlayer) PlayMove(cards []Card) error {
	log.Printf("%s attempting to play the following cards:\n", player.name)
	for _, card := range cards {
		log.Println(card.String())
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

func (player UserPlayer) RemoveCard(card Card) error {
	var remainingCards []Card
	for _, cardInHand := range *player.cards {
		if card.Suit != cardInHand.Suit || card.Value != cardInHand.Value {
			remainingCards = append(remainingCards, cardInHand)
		}
	}

	*player.cards = remainingCards

	return nil
}
