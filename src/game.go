package thirteen

import "fmt"

type Game struct {
	pile    []Card
	players [4]Player
}

func NewGame() Game {
	deck := ShuffleDeck(NewDeck())
	playerOne := NewPlayer("Player One")
	playerTwo := NewPlayer("Player Two")
	playerThree := NewPlayer("Player Three")
	playerFour := NewPlayer("Player Four")

	players := [4]Player{playerOne, playerTwo, playerThree, playerFour}

	Deal(deck, &players)

	return Game{pile: []Card{}, players: players}
}

func (game Game) ToString() (out string) {
	for _, player := range game.players {
		out = fmt.Sprintf("%sPlayer: %s\n", out, player.name)
		out = fmt.Sprintf("%sCards:\n", out)
		for _, card := range player.cards {
			out = fmt.Sprintf("%s%s\n", out, card.ToString())
		}
	}
	return out
}
