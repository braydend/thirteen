package main

import (
	"fmt"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	game := thirteen.NewGame()

	fmt.Printf("game:\n%v\n", game.ToString())
}
