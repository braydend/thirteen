package thirteen

type Game struct {
	deck Deck
	pile []Card
	// players [4]Player
}

func NewGame() Game {
	return Game{deck: NewDeck(), pile: []Card{}}
}
