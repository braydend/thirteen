package thirteen

import (
	"fmt"
	"log"
)

type Format int

const (
	CLEAR Format = iota
	SINGLE
	PAIR
	TRIPLE
	QUAD
	THREE_RUN
	FOUR_RUN
	FIVE_RUN
	SIX_RUN
	SEVEN_RUN
	EIGHT_RUN
	NINE_RUN
	TEN_RUN
	ELEVEN_RUN
	TWELVE_RUN
	THIRTEEN_RUN
	THREE_RUN_FLUSH
	FOUR_RUN_FLUSH
	FIVE_RUN_FLUSH
	SIX_RUN_FLUSH
	SEVEN_RUN_FLUSH
	EIGHT_RUN_FLUSH
	NINE_RUN_FLUSH
	TEN_RUN_FLUSH
	ELEVEN_RUN_FLUSH
	TWELVE_RUN_FLUSH
	THIRTEEN_RUN_FLUSH
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Format

// type Pile = map[int][]Card

type Game struct {
	// TODO: Change the key in this map to a player or player id
	// This will allow the game to keep track of the player with the most recent play
	// Will also need to keep track of which players have passed on the play
	// Could even add a struct to contain the format, all plays in the format, skipped players
	pile              Pile
	players           [4]Player
	activePlayerIndex int
	// currentFormat     Format
}

func NewGame() Game {
	pile := NewPile()
	pile.SetFormat(SINGLE)
	baseGame := Game{pile: pile, players: [4]Player{}}

	return baseGame
}

func (game *Game) Players() *[4]Player {
	players := game.players
	return &players
}

func (game *Game) Start() error {
	if game.PlayerCount() != 4 {
		return fmt.Errorf("Game requires 4 players to start and there are only %d.", len(game.players))
	}

	deck := ShuffleDeck(NewDeck())
	Deal(deck, game.players)

	for index, player := range game.players {
		log.Printf("player: %s\n%v", player.Name(), player.Cards())
		if player.HasCard(SPADE, THREE) {
			log.Printf("%s is active", player.Name())
			game.SetActivePlayerIndex(index)
		}
	}

	err := game.startGameLoop()

	if err != nil {
		return err
	}

	return nil
}

func (game *Game) resetFormat(skippedPlayers *map[string]Player) {
	*skippedPlayers = make(map[string]Player)
	game.pile.SetFormat(CLEAR)
	winningPlayer := game.pile.LatestPlay().player
	var winningPlayerIndex int
	for i, player := range game.players {
		if player.Id() == winningPlayer.Id() {
			winningPlayerIndex = i
		}
	}
	game.SetActivePlayerIndex(winningPlayerIndex)
}

func (game *Game) startGameLoop() error {
	skippedPlayers := make(map[string]Player)
	finishedPlayers := make(map[string]Player)
	for len(finishedPlayers) < 1 {
		if len(skippedPlayers) == 3 {
			game.resetFormat(&skippedPlayers)
		}
		log.Printf("Current format: %s\n", game.pile.Format())

		activePlayer := game.ActivePlayer()
		log.Printf("Active Player: %s\n", activePlayer.name)

		_, isSkipped := skippedPlayers[activePlayer.id]
		_, isFinished := finishedPlayers[activePlayer.id]

		if !isFinished && !isSkipped {
			canPlay, err := activePlayer.Play()
			if activePlayer.CardCount() == 0 {
				finishedPlayers[activePlayer.id] = activePlayer
			}

			if !canPlay {
				log.Printf("%s cant play: %s.", activePlayer.Name(), err)
				skippedPlayers[activePlayer.id] = activePlayer
			}
		}

		game.selectNextActivePlayer()
	}

	var winner Player

	for _, v := range finishedPlayers {
		winner = v
	}

	log.Printf("%s has won!", winner.name)

	return nil
}

func (game *Game) AddPlayer(player Player) error {
	currentPlayerCount := game.PlayerCount()

	if currentPlayerCount == 4 {
		return fmt.Errorf("Cannot have more than 4 players.")
	}

	game.players[currentPlayerCount] = player

	return nil
}

func (game *Game) SetActivePlayerIndex(playerIndex int) {
	game.activePlayerIndex = playerIndex
}

func (game Game) ActivePlayer() Player {
	return game.players[game.activePlayerIndex]
}

func (game Game) PlayerCount() int {
	count := 0
	for _, player := range game.players {
		if len(player.id) > 0 {
			count += 1
		}
	}

	return count
}

func (game *Game) Pile() *Pile {
	return &game.pile
}

func (game Game) Log() {
	log.Printf("Active Player: %s\n", game.ActivePlayer().Name())
	log.Printf("Current Format: %v\n", game.pile.Format())
	for _, player := range game.players {
		var playerLog string
		playerLog = fmt.Sprintf("%sPlayer: %s\n", playerLog, player.Name())
		playerLog = fmt.Sprintf("%sCards (%d):\n", playerLog, player.CardCount())
		for _, card := range *player.Cards() {
			playerLog = fmt.Sprintf("%s%s\n", playerLog, card.String())
		}

		log.Println(playerLog)
	}
}

func (game *Game) playMove(cards []Card, format Format, player *Player) error {
	log.Printf("Playing move with cards:\n%s\n", StringifyCards(cards))

	if game.pile.Format() != CLEAR && game.pile.Format() != format {
		return fmt.Errorf("This move is not valid for the current format.")
	}

	pile := game.Pile()
	pile.AddPlay(Play{player: player, cards: cards, format: format})
	currentFormat := pile.Format()
	if format != currentFormat {
		pile.SetFormat(format)
	}

	return nil
}

func (game *Game) selectNextActivePlayer() {
	game.SetActivePlayerIndex((game.activePlayerIndex + 1) % 4)
}

func analyzePlay(cards []Card) (Format, error) {
	isRun := true
	isMatch := true
	isFlush := true
	sortedCards := SortCards(cards)
	var currentSuit Suit
	var currentValue Value

	for i, card := range sortedCards {
		if i == 0 {
			currentSuit = card.Suit
			currentValue = card.Value
			continue
		}

		if currentSuit == card.Suit {
			isFlush = isFlush && true
		}

		if currentValue == card.Value {
			isMatch = isMatch && true
		}

		if currentValue == card.Value-1 {
			isRun = isRun && true
		}
	}

	switch len(cards) {
	case 1:
		return SINGLE, nil
	case 2:
		if isMatch {
			return PAIR, nil
		}
		break
	case 3:
		if isMatch {
			return TRIPLE, nil
		}
		if isRun {
			if isFlush {
				return THREE_RUN_FLUSH, nil
			}
			return THREE_RUN, nil
		}
		break
	case 4:
		if isMatch {
			return QUAD, nil
		}
		if isRun {
			if isFlush {
				return FOUR_RUN_FLUSH, nil
			}
			return FOUR_RUN, nil
		}
		break
	case 5:
		if isRun {
			if isFlush {
				return FIVE_RUN_FLUSH, nil
			}
			return FIVE_RUN, nil
		}
		break
	case 6:
		if isRun {
			if isFlush {
				return SIX_RUN_FLUSH, nil
			}
			return SIX_RUN, nil
		}
		break
	case 7:
		if isRun {
			if isFlush {
				return SEVEN_RUN_FLUSH, nil
			}
			return SEVEN_RUN, nil
		}
		break
	case 8:
		if isRun {
			if isFlush {
				return EIGHT_RUN_FLUSH, nil
			}
			return EIGHT_RUN, nil
		}
		break
	case 9:
		if isRun {
			if isFlush {
				return NINE_RUN_FLUSH, nil
			}
			return NINE_RUN, nil
		}
		break
	case 10:
		if isRun {
			if isFlush {
				return TEN_RUN_FLUSH, nil
			}
			return TEN_RUN, nil
		}
		break
	case 11:
		if isRun {
			if isFlush {
				return ELEVEN_RUN_FLUSH, nil
			}
			return ELEVEN_RUN, nil
		}
		break
	case 12:
		if isRun {
			if isFlush {
				return TWELVE_RUN_FLUSH, nil
			}
			return TWELVE_RUN, nil
		}
		break
	case 13:
		if isRun {
			if isFlush {
				return THIRTEEN_RUN_FLUSH, nil
			}
			return THIRTEEN_RUN, nil
		}
		break
	}

	return 0, fmt.Errorf("Unable to find a play with the cards submitted.\nCards:\n%s", StringifyCards(cards))
}
