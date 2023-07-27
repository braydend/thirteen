package thirteen_test

import (
	"testing"

	thirteen "github.com/braydend/thirteen/src"
)

func TestNewGame(t *testing.T) {
	result := thirteen.NewGame()

	if result.Format() != thirteen.SINGLE {
		t.Errorf("Game should start with SINGLE format, but got %v", result.Format())
	}
}

func TestAddPlayer(t *testing.T) {
	type testInput struct {
		players []thirteen.Player
	}

	type testCase struct {
		name                string
		input               testInput
		expectError         bool
		expectedPlayerCount int
	}

	stubPlayer := thirteen.NewPlayer("Stub Player", func(cards []thirteen.Card) error { return nil })

	testCases := []testCase{
		{"Can add one player", testInput{[]thirteen.Player{stubPlayer}}, false, 1},
		{"Can add two players", testInput{
			[]thirteen.Player{stubPlayer, stubPlayer},
		}, false, 2},
		{"Can add three players", testInput{
			[]thirteen.Player{stubPlayer, stubPlayer, stubPlayer},
		}, false, 3},
		{"Can add four players", testInput{
			[]thirteen.Player{stubPlayer, stubPlayer, stubPlayer, stubPlayer},
		}, false, 4},
		{"Cannot add five players", testInput{
			[]thirteen.Player{stubPlayer, stubPlayer, stubPlayer, stubPlayer, thirteen.Player{}},
		}, true, 4},
	}

	for _, testCase := range testCases {
		game := thirteen.NewGame()

		for _, player := range testCase.input.players {
			err := game.AddPlayer(player)

			if !testCase.expectError && err != nil {
				t.Errorf("Failed to add player %v", err)
			}
		}

		if testCase.expectedPlayerCount != game.PlayerCount() {
			t.Errorf("Expected %d players, but game has %d", testCase.expectedPlayerCount, game.PlayerCount())
		}
	}
}

func TestStartErrors(t *testing.T) {
	type testInput struct {
		game thirteen.Game
	}

	type testCase struct {
		name  string
		input testInput
	}

	testCases := []testCase{
		{"Fails to start game with less than four players", testInput{thirteen.NewGame()}},
	}

	for _, testCase := range testCases {
		err := testCase.input.game.Start()

		if err == nil {
			t.Errorf("%s expected error but none was returned.", testCase.name)
		}
	}
}

func TestStart(t *testing.T) {
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
		t.Error(err)
	}
}
