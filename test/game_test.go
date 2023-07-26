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

	testCases := []testCase{
		{"Can add one player", testInput{[]thirteen.Player{thirteen.UserPlayer{}}}, false, 1},
		{"Can add two players", testInput{
			[]thirteen.Player{thirteen.UserPlayer{}, thirteen.UserPlayer{}},
		}, false, 2},
		{"Can add three players", testInput{
			[]thirteen.Player{thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}},
		}, false, 3},
		{"Can add four players", testInput{
			[]thirteen.Player{thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}},
		}, false, 4},
		{"Cannot add five players", testInput{
			[]thirteen.Player{thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}, thirteen.UserPlayer{}},
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

}
