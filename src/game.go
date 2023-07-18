package thirteen

import (
	"fmt"
	"log"
)

type Format = uint8

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

type Game struct {
	pile          *map[int][]Card
	players       *[4]Player
	activePlayer  Player
	currentFormat Format
}

func NewGame() Game {
	baseGame := Game{pile: &map[int][]Card{}}
	deck := ShuffleDeck(NewDeck())
	playerOne := NewPlayer("Player One", baseGame.PlayMove)
	playerTwo := NewPlayer("Player Two", baseGame.PlayMove)
	playerThree := NewPlayer("Player Three", baseGame.PlayMove)
	playerFour := NewPlayer("Player Four", baseGame.PlayMove)

	players := [4]Player{playerOne, playerTwo, playerThree, playerFour}

	Deal(deck, &players)

	baseGame.players = &players

	var activePlayer Player

	for _, player := range players {
		if player.HasCard(SPADE, THREE) {
			activePlayer = player
		}
	}

	baseGame.activePlayer = activePlayer

	return baseGame
}

func (game *Game) ActivePlayer() *Player {
	return &game.activePlayer
}

func (game *Game) Pile() *map[int][]Card {
	return game.pile
}

func (game Game) Log() {
	log.Printf("Active Player: %s\n", game.activePlayer.name)
	log.Printf("Current Format: %v\n", game.currentFormat)
	log.Printf("Pile has:\n%v\n", *game.pile)
	for _, player := range game.players {
		var playerLog string
		playerLog = fmt.Sprintf("%sPlayer: %s\n", playerLog, player.name)
		playerLog = fmt.Sprintf("%sCards:\n", playerLog)
		for _, card := range *player.cards {
			playerLog = fmt.Sprintf("%s%s\n", playerLog, card.ToString())
		}

		log.Println(playerLog)
	}
}

func (game *Game) Format() Format {
	return game.currentFormat
}

func (game *Game) SetFormat(format Format) {
	game.currentFormat = format
}

func (game *Game) AddToPile(cards []Card) {
	playsInPile := len(*game.pile)
	currentPile := *game.pile
	currentPile[playsInPile] = cards

	game.pile = &currentPile
}

func (game *Game) PlayMove(cards []Card) error {
	log.Printf("Playing move with cards:\n%s\n", StringifyCards(cards))
	format, err := analyzePlay(cards)

	if err != nil {
		return err
	}

	game.SetFormat(format)
	game.AddToPile(cards)

	return nil
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
