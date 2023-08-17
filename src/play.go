package thirteen

import "fmt"

type Play struct {
	player *Player
	cards  []Card
	format Format
}

func NewPlay() Play {
	return Play{}
}

func (play Play) String() string {
	return fmt.Sprintf("%s - %s: %s", play.format, play.player.Name(), StringifyCards(play.cards))
}
