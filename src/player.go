package thirteen

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type GameData struct {
	format *Format
	pile   *Pile
}

type Player struct {
	id       string
	name     string
	cards    *[]Card
	game     *Game
	isCpu    bool
	gameData GameData
}

func NewPlayer(name string, game *Game) Player {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatalf("Failed to generate UUID for player. %s", err)
	}

	return Player{id: uuid.String(), name: name, cards: &[]Card{}, game: game}
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

func (player *Player) Play() (bool, error) {
	if player.isCpu {
		currentFormat := player.game.pile.Format()
		play, playedFormat, err := AutoPlay(*player.Cards(), &currentFormat, *player.game.Pile())

		if len(play) == 0 {
			return false, fmt.Errorf("No plays available")
		}

		if err != nil {
			return false, err
		}

		err = player.game.playMove(play, playedFormat, player)

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
