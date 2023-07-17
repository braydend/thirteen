package main

import (
	"fmt"

	thirteen "github.com/braydend/thirteen/src"
)

func main() {
	deck := thirteen.NewDeck()

	fmt.Printf("sorted deck:\n%v\n", thirteen.StringifyDeck(thirteen.SortDeck(deck)))
	fmt.Printf("shuffled deck:\n%v\n", thirteen.StringifyDeck(thirteen.ShuffleDeck(deck)))
}
