package thirteen

type Player interface {
	AddCard(card *Card)
	CardCount() uint8
	Id() string
	Name() string
	Cards() *[]Card
	Play() (bool, error)
	HasCard(suit Suit, value Value) bool
	RemoveCard(card Card) error
}

func CastToPlayer[P Player](player P) Player {
	return player
}

func CastToPlayers[P Player](players [4]P) (out [4]Player) {
	for i, player := range players {
		out[i] = player
	}

	return out
}
