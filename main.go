package main

import (
	"fmt"

	card "github.com/braydend/thirteen/src"
)

func main() {
	deck := card.NewDeck()

	fmt.Printf("sorted deck:\n%v\n", card.StringifyDeck(card.SortDeck(deck)))
	fmt.Printf("shuffled deck:\n%v\n", card.StringifyDeck(card.ShuffleDeck(deck)))
}
