package main

import (
	"log"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	game := thirteen.NewGame()
	playerOne := thirteen.NewPlayer("P1", game.PlayMove)
	playerOne.SetCPU(true)
	playerTwo := thirteen.NewPlayer("P2", game.PlayMove)
	playerTwo.SetCPU(true)
	playerThree := thirteen.NewPlayer("P3", game.PlayMove)
	playerThree.SetCPU(true)
	playerFour := thirteen.NewPlayer("P4", game.PlayMove)
	playerFour.SetCPU(true)
	game.AddPlayer(playerOne)
	game.AddPlayer(playerTwo)
	game.AddPlayer(playerThree)
	game.AddPlayer(playerFour)

	err := game.Start()

	if err != nil {
		log.Fatal(err)
	}

	// player := game.ActivePlayer()
	// err := player.Player().PlayMove([]thirteen.Card{thirteen.NewCard(thirteen.THREE, thirteen.SPADE)})

	// canPlay, err := (*game.ActivePlayer()).Play()

	// if err != nil {
	// 	log.Fatalf("Failed to play move. %s", err)
	// }

	// game.SetFormat(thirteen.SINGLE)

	// game.Log()

	// _, err = (*game.ActivePlayer()).Play()

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("Can play: %v\n", canPlay)

	// game.Log()
}
