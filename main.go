package main

import (
	"log"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	game := thirteen.NewGame()

	game.Log()

	// player := game.ActivePlayer()
	// err := player.Player().PlayMove([]thirteen.Card{thirteen.NewCard(thirteen.THREE, thirteen.SPADE)})

	canPlay, err := (*game.ActivePlayer()).Play()

	if err != nil {
		log.Fatalf("Failed to play move. %s", err)
	}

	game.SetFormat(thirteen.SINGLE)

	game.Log()

	_, err = (*game.ActivePlayer()).Play()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Can play: %v\n", canPlay)

	game.Log()
}
