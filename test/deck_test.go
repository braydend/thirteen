package thirteen_test

import (
	"reflect"
	"testing"

	thirteen "github.com/braydend/thirteen/src"
)

func TestNewDeck(t *testing.T) {
	expected := []thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.THREE},
		{Suit: thirteen.SPADE, Value: thirteen.FOUR},
		{Suit: thirteen.SPADE, Value: thirteen.FIVE},
		{Suit: thirteen.SPADE, Value: thirteen.SIX},
		{Suit: thirteen.SPADE, Value: thirteen.SEVEN},
		{Suit: thirteen.SPADE, Value: thirteen.EIGHT},
		{Suit: thirteen.SPADE, Value: thirteen.NINE},
		{Suit: thirteen.SPADE, Value: thirteen.TEN},
		{Suit: thirteen.SPADE, Value: thirteen.JACK},
		{Suit: thirteen.SPADE, Value: thirteen.QUEEN},
		{Suit: thirteen.SPADE, Value: thirteen.KING},
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.SPADE, Value: thirteen.TWO},
		{Suit: thirteen.CLUB, Value: thirteen.THREE},
		{Suit: thirteen.CLUB, Value: thirteen.FOUR},
		{Suit: thirteen.CLUB, Value: thirteen.FIVE},
		{Suit: thirteen.CLUB, Value: thirteen.SIX},
		{Suit: thirteen.CLUB, Value: thirteen.SEVEN},
		{Suit: thirteen.CLUB, Value: thirteen.EIGHT},
		{Suit: thirteen.CLUB, Value: thirteen.NINE},
		{Suit: thirteen.CLUB, Value: thirteen.TEN},
		{Suit: thirteen.CLUB, Value: thirteen.JACK},
		{Suit: thirteen.CLUB, Value: thirteen.QUEEN},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.CLUB, Value: thirteen.ACE},
		{Suit: thirteen.CLUB, Value: thirteen.TWO},
		{Suit: thirteen.DIAMOND, Value: thirteen.THREE},
		{Suit: thirteen.DIAMOND, Value: thirteen.FOUR},
		{Suit: thirteen.DIAMOND, Value: thirteen.FIVE},
		{Suit: thirteen.DIAMOND, Value: thirteen.SIX},
		{Suit: thirteen.DIAMOND, Value: thirteen.SEVEN},
		{Suit: thirteen.DIAMOND, Value: thirteen.EIGHT},
		{Suit: thirteen.DIAMOND, Value: thirteen.NINE},
		{Suit: thirteen.DIAMOND, Value: thirteen.TEN},
		{Suit: thirteen.DIAMOND, Value: thirteen.JACK},
		{Suit: thirteen.DIAMOND, Value: thirteen.QUEEN},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.DIAMOND, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.TWO},
		{Suit: thirteen.HEART, Value: thirteen.THREE},
		{Suit: thirteen.HEART, Value: thirteen.FOUR},
		{Suit: thirteen.HEART, Value: thirteen.FIVE},
		{Suit: thirteen.HEART, Value: thirteen.SIX},
		{Suit: thirteen.HEART, Value: thirteen.SEVEN},
		{Suit: thirteen.HEART, Value: thirteen.EIGHT},
		{Suit: thirteen.HEART, Value: thirteen.NINE},
		{Suit: thirteen.HEART, Value: thirteen.TEN},
		{Suit: thirteen.HEART, Value: thirteen.JACK},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
		{Suit: thirteen.HEART, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.HEART, Value: thirteen.TWO},
	}

	result := thirteen.NewDeck()

	isEqual := reflect.DeepEqual(thirteen.SortCards(result), thirteen.SortCards(expected))

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expected)
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := thirteen.NewDeck()
	shuffledDeck := thirteen.ShuffleDeck(thirteen.NewDeck())

	isEqual := reflect.DeepEqual(deck, shuffledDeck)

	if isEqual {
		t.Errorf("%v should not equal %v", deck, shuffledDeck)
	}
}

func TestDeal(t *testing.T) {
	stubPlayFn := func(_ []thirteen.Card) error {
		return nil
	}
	players := [4]thirteen.UserPlayer{
		thirteen.NewUserPlayer("Stub Player One", stubPlayFn),
		thirteen.NewUserPlayer("Stub Player Two", stubPlayFn),
		thirteen.NewUserPlayer("Stub Player Three", stubPlayFn),
		thirteen.NewUserPlayer("Stub Player Four", stubPlayFn),
	}

	deck := thirteen.NewDeck()

	thirteen.Deal(deck, players)

	expectedCardCount := uint8(13)

	for _, player := range players {
		isEqual := player.CardCount() == expectedCardCount

		if !isEqual {
			t.Errorf("%v does not equal the expected %v", player.CardCount(), expectedCardCount)
		}
	}
}
