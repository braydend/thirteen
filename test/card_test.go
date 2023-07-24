package thirteen_test

import (
	"reflect"
	"testing"

	thirteen "github.com/braydend/thirteen/src"
)

func TestNewCard(t *testing.T) {
	// Valid cards
	for _, value := range []thirteen.Value{thirteen.THREE, thirteen.FOUR, thirteen.FIVE, thirteen.SIX, thirteen.SEVEN, thirteen.EIGHT, thirteen.NINE, thirteen.TEN, thirteen.JACK, thirteen.QUEEN, thirteen.KING, thirteen.ACE, thirteen.TWO} {
		for _, suit := range []thirteen.Suit{thirteen.SPADE, thirteen.CLUB, thirteen.DIAMOND, thirteen.HEART} {
			result, _ := thirteen.NewCard(value, suit)
			expected := thirteen.Card{Suit: suit, Value: value}

			isEqual := reflect.DeepEqual(result, expected)

			if !isEqual {
				t.Errorf("%v does not equal the expected %v", result, expected)
			}
		}
	}

	//Invalid Cards
	for _, value := range []uint8{0, 1, 2, 16, 17} {
		for _, suit := range []uint8{0, 5, 6, 7} {
			_, err := thirteen.NewCard(value, suit)

			if err == nil {
				t.Errorf("Expected an error when creating a card with the value %v and suit %v", value, suit)
			}
		}
	}
}

func TestString(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.THREE},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.SPADE, Value: thirteen.SEVEN},
	}
	for i, card := range cards {
		var expected string
		result := card.String()
		switch i {
		case 0:
			expected = "Ace of Hearts"
		case 1:
			expected = "Three of Diamonds"
		case 2:
			expected = "King of Clubs"
		case 3:
			expected = "Seven of Spades"
		}

		isEqual := reflect.DeepEqual(result, expected)

		if !isEqual {
			t.Errorf("%v does not equal the expected %v", result, expected)
		}
	}
}

func TestSortCards(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.CLUB, Value: thirteen.THREE},
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.SPADE, Value: thirteen.SEVEN},
		{Suit: thirteen.DIAMOND, Value: thirteen.THREE},
		{Suit: thirteen.SPADE, Value: thirteen.THREE},
	}

	result := thirteen.SortCards(cards)
	expected := []thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.THREE},
		{Suit: thirteen.CLUB, Value: thirteen.THREE},
		{Suit: thirteen.DIAMOND, Value: thirteen.THREE},
		{Suit: thirteen.SPADE, Value: thirteen.SEVEN},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.ACE},
	}

	isEqual := reflect.DeepEqual(result, expected)

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expected)
	}
}
