package main

import (
	"fmt"

	card "github.com/braydend/thirteen/src"
)

func main() {
	newCard := card.NewCard(card.THREE, card.DIAMOND)

	fmt.Printf("card: %v", newCard)
}
