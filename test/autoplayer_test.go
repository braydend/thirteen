package thirteen_test

import (
	"reflect"
	"testing"

	thirteen "github.com/braydend/thirteen/src"
)

func buildAutoPlayer(cards []thirteen.Card, game thirteen.Game) thirteen.AutoPlayer {
	stubPlayFn := func(_ []thirteen.Card) error {
		return nil
	}
	format := game.Format()
	pile := game.Pile()

	userPlayer := thirteen.NewUserPlayer("Stub Player", stubPlayFn)

	player := thirteen.NewAutoPlayer(&userPlayer, &format, pile)

	for _, card := range cards {
		player.AddCard(&card)
	}

	return player
}

func TestNewAutoPlayer(t *testing.T) {
	stubPlayFn := func(_ []thirteen.Card) error {
		return nil
	}
	format := thirteen.SINGLE
	pile := map[int][]thirteen.Card{}

	userPlayer := thirteen.NewUserPlayer("Stub Player", stubPlayFn)

	expected := thirteen.AutoPlayer{BasePlayer: &userPlayer, CurrentFormat: &format, CurrentPile: &pile}
	result := thirteen.NewAutoPlayer(&userPlayer, &format, &pile)

	isEqual := reflect.DeepEqual(result, expected)

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expected)
	}
}

func TestAddCard(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
	}

	game := thirteen.NewGame()

	player := buildAutoPlayer(cards, game)

	expectedCount := uint8(3)
	expectedCards := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
	}

	// Assert count is correct

	isEqualCount := reflect.DeepEqual(player.CardCount(), expectedCount)

	if !isEqualCount {
		t.Errorf("%v does not equal the expected %v", player.CardCount(), expectedCount)
	}

	// Assert cards is in player's hand are correct

	isEqualCards := reflect.DeepEqual(*player.Cards(), expectedCards)

	if !isEqualCards {
		t.Errorf("%v does not equal the expected %v", *player.Cards(), expectedCards)
	}

	// Assert card can be correctly identified in hand

	if !player.HasCard(thirteen.HEART, thirteen.QUEEN) {
		t.Errorf("%v does not have the card %v", player.Name(), thirteen.Card{Suit: thirteen.HEART, Value: thirteen.QUEEN})
	}
}

func TestRemoveCard(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
	}

	game := thirteen.NewGame()

	player := buildAutoPlayer(cards, game)

	expectedCount := uint8(1)
	expectedCards := []thirteen.Card{
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
	}

	player.RemoveCard(thirteen.Card{Suit: thirteen.HEART, Value: thirteen.QUEEN})
	player.RemoveCard(thirteen.Card{Suit: thirteen.SPADE, Value: thirteen.ACE})

	// Assert count is correct

	isEqualCount := reflect.DeepEqual(player.CardCount(), expectedCount)

	if !isEqualCount {
		t.Errorf("%v does not equal the expected %v", player.CardCount(), expectedCount)
	}

	// Assert cards is in player's hand are correct

	isEqualCards := reflect.DeepEqual(*player.Cards(), expectedCards)

	if !isEqualCards {
		t.Errorf("%v does not equal the expected %v", *player.Cards(), expectedCards)
	}

	// Assert card can be correctly identified in hand

	if !player.HasCard(thirteen.DIAMOND, thirteen.KING) {
		t.Errorf("%v does not have the card %v", player.Name(), thirteen.Card{Suit: thirteen.DIAMOND, Value: thirteen.KING})
	}
}

func TestBuildPlaySingle(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}
	game := thirteen.NewGame()
	game.SetFormat(thirteen.SINGLE)
	player := buildAutoPlayer(cards, game)
	expectedPlay := []thirteen.Card{
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}
	result := player.BuildPlay()

	isEqual := reflect.DeepEqual(result, expectedPlay)

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expectedPlay)
	}
}

func TestBuildPlayPair(t *testing.T) {
	cards := []thirteen.Card{
		{Suit: thirteen.SPADE, Value: thirteen.ACE},
		{Suit: thirteen.CLUB, Value: thirteen.ACE},
		{Suit: thirteen.HEART, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.ACE},
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.SPADE, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.KING},
		{Suit: thirteen.CLUB, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.QUEEN},
	}
	game := thirteen.NewGame()
	game.SetFormat(thirteen.PAIR)
	player := buildAutoPlayer(cards, game)
	expectedPlay := []thirteen.Card{
		{Suit: thirteen.DIAMOND, Value: thirteen.KING},
		{Suit: thirteen.HEART, Value: thirteen.KING},
	}
	result := player.BuildPlay()

	isEqual := reflect.DeepEqual(result, expectedPlay)

	if !isEqual {
		t.Errorf("%v does not equal the expected %v", result, expectedPlay)
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
		game := thirteen.NewGame()
		game.SetFormat(testCase.input.format)
		player := buildAutoPlayer(testCase.input.cards, game)
		result := player.BuildPlay()

		isEqual := reflect.DeepEqual(result, testCase.expected)

		if !isEqual {
			t.Errorf("%s failed. %v does not equal the expected %v", testCase.label, thirteen.StringifyCards(result), thirteen.StringifyCards(testCase.expected))
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
		game := thirteen.NewGame()
		game.SetFormat(testCase.input.format)
		player := buildAutoPlayer(testCase.input.cards, game)
		result := player.BuildPlay()

		isEqual := reflect.DeepEqual(result, testCase.expected)

		if !isEqual {
			t.Errorf("%s failed. %v does not equal the expected %v", testCase.label, thirteen.StringifyCards(result), thirteen.StringifyCards(testCase.expected))
		}

	}
}
