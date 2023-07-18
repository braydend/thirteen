package main

import (
	"fmt"
	"log"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	game := thirteen.NewGame()

	game.Log()

	player := game.ActivePlayer()
	err := player.PlayMove([]thirteen.Card{thirteen.NewCard(thirteen.THREE, thirteen.SPADE)})

	if err != nil {
		log.Fatalf("Failed to play move. %s", err)
	}

	pile := game.Pile()
	for _, play := range *pile {
		fmt.Printf("---START PLAY---\n%s---END PLAY---\n", thirteen.StringifyCards(play))
	}

	game.Log()
}
