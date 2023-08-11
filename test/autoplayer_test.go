package thirteen_test

import (
	"reflect"
	"testing"

	thirteen "github.com/braydend/thirteen/src"
)

func TestBuildPlaySingle(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}
	expectedPlay := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}
	currentFormat := thirteen.SINGLE
	result, format, _ := thirteen.AutoPlay(cards, &currentFormat, thirteen.Pile{})

	isEqual := reflect.DeepEqual(result, expectedPlay)

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expectedPlay)
	}

	if format != thirteen.SINGLE {
		t.Errorf("%v is not a single card", result)
	}
}

func TestBuildPlayMatch(t *testing.T) {
	cards := &[]thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.CLUB, Value: thirteen.ACE},
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.ACE},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.SPADE, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.KING},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}

	type input struct {
		cards  []thirteen.Card
		format thirteen.Format
	}

	type testCase struct {
		label    string
		input    input
		expected []thirteen.Card
	}

	testCases := []testCase{
		{"Single card", input{*cards, thirteen.SINGLE}, []thirteen.Card{{Suit: thirteen.HEART, Value: thirteen.QUEEN}}},
		{"Pair", input{*cards, thirteen.PAIR}, []thirteen.Card{
			{Suit: thirteen.SPADE, Value: thirteen.KING},
			{Suit: thirteen.CLUB, Value: thirteen.KING},
		}},
		{"Triple", input{*cards, thirteen.TRIPLE}, []thirteen.Card{
			{Suit: thirteen.SPADE, Value: thirteen.KING},
			{Suit: thirteen.CLUB, Value: thirteen.KING},
			{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		}},
		{"Quad", input{*cards, thirteen.QUAD}, []thirteen.Card{
			{Suit: thirteen.SPADE, Value: thirteen.KING},
			{Suit: thirteen.CLUB, Value: thirteen.KING},
			{Suit: thirteen.DIAMOND, Value: thirteen.KING},
			{Suit: thirteen.HEART, Value: thirteen.KING},
		}},
	}

	for _, testCase := range testCases {
		currentFormat := testCase.input.format
		result, format, _ := thirteen.AutoPlay(testCase.input.cards, &currentFormat, thirteen.Pile{})

		isEqual := reflect.DeepEqual(result, testCase.expected)

		if !isEqual {
			t.Errorf("%s failed. %v does not equal the expected %v", testCase.label, thirteen.StringifyCards(result), thirteen.StringifyCards(testCase.expected))
		}
		if format != testCase.input.format {
			t.Errorf("%s failed. %v is not the correct format: %v", testCase.label, thirteen.StringifyCards(result), testCase.input.format)
		}
	}
}

func TestBuildPlayRun(t *testing.T) {
	cards := &[]thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.ACE},
		{Suit: thirteen.HEART, Value: thirteen.KING},
		{Suit: thirteen.DIAMOND, Value: thirteen.SEVEN},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
		{Suit: thirteen.HEART, Value: thirteen.TWO},
		{Suit: thirteen.DIAMOND, Value: thirteen.THREE},
		{Suit: thirteen.CLUB, Value: thirteen.JACK},
		{Suit: thirteen.SPADE, Value: thirteen.TEN},
		{Suit: thirteen.DIAMOND, Value: thirteen.NINE},
		{Suit: thirteen.DIAMOND, Value: thirteen.FOUR},
	}

	type input struct {
		cards  []thirteen.Card
		format thirteen.Format
	}

	type testCase struct {
		label    string
		input    input
		expected []thirteen.Card
	}

	testCases := []testCase{
		{"Three card run", input{*cards, thirteen.THREE_RUN}, []thirteen.Card{
			{Suit: thirteen.DIAMOND, Value: thirteen.NINE},
			{Suit: thirteen.SPADE, Value: thirteen.TEN},
			{Suit: thirteen.CLUB, Value: thirteen.JACK},
		}},
		{"Three card flush run", input{*cards, thirteen.THREE_RUN_FLUSH}, []thirteen.Card{
			{Suit: thirteen.HEART, Value: thirteen.QUEEN},
			{Suit: thirteen.HEART, Value: thirteen.KING},
			{Suit: thirteen.HEART, Value: thirteen.ACE},
		}},
		{"Five card run", input{*cards, thirteen.FIVE_RUN}, []thirteen.Card{
			{Suit: thirteen.DIAMOND, Value: thirteen.NINE},
			{Suit: thirteen.SPADE, Value: thirteen.TEN},
			{Suit: thirteen.CLUB, Value: thirteen.JACK},
			{Suit: thirteen.HEART, Value: thirteen.QUEEN},
			{Suit: thirteen.HEART, Value: thirteen.KING},
		}},
		{"Seven card run", input{*cards, thirteen.SEVEN_RUN}, []thirteen.Card{
			{Suit: thirteen.DIAMOND, Value: thirteen.NINE},
			{Suit: thirteen.SPADE, Value: thirteen.TEN},
			{Suit: thirteen.CLUB, Value: thirteen.JACK},
			{Suit: thirteen.HEART, Value: thirteen.QUEEN},
			{Suit: thirteen.HEART, Value: thirteen.KING},
			{Suit: thirteen.DIAMOND, Value: thirteen.ACE},
			{Suit: thirteen.HEART, Value: thirteen.TWO},
		}},
	}

	for _, testCase := range testCases {
		currentFormat := testCase.input.format
		result, format, _ := thirteen.AutoPlay(testCase.input.cards, &currentFormat, thirteen.Pile{})

		isEqual := reflect.DeepEqual(result, testCase.expected)

		if !isEqual {
			t.Errorf("%s failed. %v does not equal the expected %v", testCase.label, thirteen.StringifyCards(result), thirteen.StringifyCards(testCase.expected))
		}

		if format != testCase.input.format {
			t.Errorf("%s failed. %v is not the correct format: %v", testCase.label, thirteen.StringifyCards(result), testCase.input.format)
		}
	}
}
