package thirteen

import "fmt"

type Game struct {
	pile         []Card
	players      [4]Player
	activePlayer Player
}

func NewGame() Game {
	deck := ShuffleDeck(NewDeck())
	playerOne := NewPlayer("Player One")
	playerTwo := NewPlayer("Player Two")
	playerThree := NewPlayer("Player Three")
	playerFour := NewPlayer("Player Four")

	players := [4]Player{playerOne, playerTwo, playerThree, playerFour}

	Deal(deck, &players)

	var activePlayer Player

	for _, player := range players {
		if player.HasCard(SPADE, THREE) {
			activePlayer = player
		}
	}

	return Game{pile: []Card{}, players: players, activePlayer: activePlayer}
}

func (game Game) ToString() (out string) {
	out = fmt.Sprintf("%sActive Player: %s\n", out, game.activePlayer.name)
	for _, player := range game.players {
		out = fmt.Sprintf("%sPlayer: %s\n", out, player.name)
		out = fmt.Sprintf("%sCards:\n", out)
		for _, card := range player.cards {
			out = fmt.Sprintf("%s%s\n", out, card.ToString())
		}
	}
	return out
}
