package thirteen

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

// type Player interface {
// 	AddCard(card *Card)
// 	CardCount() uint8
// 	Id() string
// 	Name() string
// 	Cards() *[]Card
// 	Play() (bool, error)
// 	HasCard(suit Suit, value Value) bool
// 	RemoveCard(card Card) error
// }

type GameData struct {
	format *Format
	pile   *Pile
}

type Player struct {
	id       string
	name     string
	cards    *[]Card
	playMove func(hand []Card) error
	isCpu    bool
	gameData GameData
}

// TODO: Add fn for SetPlayFn() to Player interface
// playFn is currently tied into user player but needs to be used across all
// implementations of Player interface
func NewPlayer(name string, playFn func(cards []Card) error) Player {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatalf("Failed to generate UUID for player. %s", err)
	}

	return Player{id: uuid.String(), name: name, cards: &[]Card{}, playMove: playFn}
}

func (player Player) Id() string {
	return player.id
}

func (player Player) Name() string {
	return player.name
}

func (player Player) Cards() *[]Card {
	return player.cards
}

func (player *Player) SetCPU(isCpu bool) {
	player.isCpu = isCpu
}

func (player Player) AddCard(card *Card) {
	currentCards := *player.cards
	currentCards = append(currentCards, *card)

	*player.cards = currentCards
}

func (player Player) SetGameData(data GameData) {
	player.gameData = data
}

func (player Player) CardCount() uint8 {
	return uint8(len(*player.cards))
}

func (player Player) Play(format Format, pile Pile) (bool, error) {
	if player.isCpu {
		play, err := AutoPlay(*player.Cards(), format, pile)

		if err != nil {
			return false, err
		}

		err = player.playMove(play)

		if err != nil {
			return false, err
		}

		for _, currentCard := range play {
			player.RemoveCard(currentCard)
		}

		return true, nil
	}

	// TODO: Implement this
	return false, fmt.Errorf("NOT IMPLEMENTED")
}

func (player Player) Validate() error {
	if player.gameData.format == nil {
		return fmt.Errorf("No reference to game format")
	}

	if player.gameData.pile == nil {
		return fmt.Errorf("No reference to game pile")
	}

	return nil
}

func (player Player) HasCard(suit Suit, value Value) bool {
	for _, card := range *player.cards {
		if card.Suit == suit && card.Value == value {
			return true
		}
	}

	return false
}

// func (player *Player) PlayMove(cards []Card) error {
// 	log.Printf("%s attempting to play the following cards:\n", player.name)
// 	for _, card := range cards {
// 		log.Println(card.String())
// 	}

// 	err := player.playMove(cards)

// 	if err != nil {
// 		return err
// 	}

// 	for _, currentCard := range cards {
// 		player.RemoveCard(currentCard)
// 	}

// 	return nil
// }

func (player Player) RemoveCard(card Card) error {
	var remainingCards []Card
	for _, cardInHand := range *player.cards {
		if card.Suit != cardInHand.Suit || card.Value != cardInHand.Value {
			remainingCards = append(remainingCards, cardInHand)
		}
	}

	*player.cards = remainingCards

	return nil
}
