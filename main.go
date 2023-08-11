package main

import (
	"log"
	"os"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln("Failed to create log file.")
	}
	log.SetOutput(logFile)

	game := thirteen.NewGame()
	playerOne := thirteen.NewPlayer("P1", &game)
	playerOne.SetCPU(true)
	playerTwo := thirteen.NewPlayer("P2", &game)
	playerTwo.SetCPU(true)
	playerThree := thirteen.NewPlayer("P3", &game)
	playerThree.SetCPU(true)
	playerFour := thirteen.NewPlayer("P4", &game)
	playerFour.SetCPU(true)
	game.AddPlayer(playerOne)
	game.AddPlayer(playerTwo)
	game.AddPlayer(playerThree)
	game.AddPlayer(playerFour)

	err = game.Start()

	if err != nil {
		log.Fatal(err)
	}
}
