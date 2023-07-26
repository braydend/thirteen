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
	// TODO: Change the key in this map to a player or player id
	// This will allow the game to keep track of the player with the most recent play
	// Will also need to keep track of which players have passed on the play
	pile          *map[int][]Card
	players       *[4]Player
	activePlayer  *Player
	currentFormat Format
}

func NewGame() Game {
	baseGame := Game{pile: &map[int][]Card{}, currentFormat: SINGLE}
	// deck := ShuffleDeck(NewDeck())
	// playerOne := NewUserPlayer("Player One", baseGame.PlayMove)
	// playerTwo := NewUserPlayer("Player Two", baseGame.PlayMove)
	// playerThree := NewUserPlayer("Player Three", baseGame.PlayMove)
	// playerFour := NewUserPlayer("Player Four", baseGame.PlayMove)

	// autoOne := NewAutoPlayer(&playerOne, &baseGame.currentFormat, baseGame.pile)
	// autoTwo := NewAutoPlayer(&playerTwo, &baseGame.currentFormat, baseGame.pile)
	// autoThree := NewAutoPlayer(&playerThree, &baseGame.currentFormat, baseGame.pile)
	// autoFour := NewAutoPlayer(&playerFour, &baseGame.currentFormat, baseGame.pile)

	// autoPlayers := [4]AutoPlayer{autoOne, autoTwo, autoThree, autoFour}

	// Deal(deck, autoPlayers)

	// castedPlayers := CastToPlayers(autoPlayers)

	// baseGame.players = &castedPlayers
	baseGame.players = &[4]Player{}

	// var activePlayer AutoPlayer

	// for _, player := range autoPlayers {
	// 	log.Printf("player: %s\n%v", player.BasePlayer.name, player.BasePlayer.cards)
	// 	if player.BasePlayer.HasCard(SPADE, THREE) {
	// 		log.Printf("%s is active", player.BasePlayer.name)
	// 		activePlayer = player
	// 	}
	// }

	// castedPlayer := CastToPlayer(activePlayer)

	// baseGame.activePlayer = &castedPlayer

	return baseGame
}

func (game Game) Start() error {
	if game.PlayerCount() != 4 {
		return fmt.Errorf("Game requires 4 players to start and there are only %d.", len(game.players))
	}

	deck := ShuffleDeck(NewDeck())
	castedPlayers := CastToPlayers(*game.players)
	Deal(deck, castedPlayers)

	var activePlayer Player

	for _, player := range castedPlayers {
		log.Printf("player: %s\n%v", player.Name(), player.Cards())
		if player.HasCard(SPADE, THREE) {
			log.Printf("%s is active", player.Name())
			activePlayer = player
		}
	}

	castedPlayer := CastToPlayer(activePlayer)

	game.activePlayer = &castedPlayer

	return nil
}

func (game Game) AddPlayer(player Player) error {
	currentPlayerCount := game.PlayerCount()

	if currentPlayerCount == 4 {
		return fmt.Errorf("Cannot have more than 4 players.")
	}

	game.players[currentPlayerCount] = player

	return nil
}

func (game Game) ActivePlayer() *Player {
	return game.activePlayer
}

func (game Game) PlayerCount() int {
	count := 0
	for _, player := range *game.players {
		if player != nil {
			count += 1
		}
	}

	return count
}

func (game *Game) Pile() *map[int][]Card {
	return game.pile
}

func (game Game) Log() {
	log.Printf("Active Player: %s\n", (*game.activePlayer).Name())
	log.Printf("Current Format: %v\n", game.currentFormat)
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

	log.Println("---PILE START---")
	for _, play := range *game.pile {
		log.Println(StringifyCards(play))
	}
	log.Println("---PILE END---")
}

func (game *Game) PlayMove(cards []Card) error {
	log.Printf("Playing move with cards:\n%s\n", StringifyCards(cards))
	format, err := analyzePlay(cards)

	if err != nil {
		return err
	}

	if game.currentFormat != CLEAR && game.currentFormat != format {
		return fmt.Errorf("This move is not valid for the current format.")
	}

	game.SetFormat(format)
	game.AddToPile(cards)
	game.SelectNextActivePlayer()

	return nil
}

func (game *Game) SelectNextActivePlayer() {
	var activePlayerIndex int
	for i, player := range game.players {
		activePlayer := *game.activePlayer
		if player.Id() == activePlayer.Id() {
			activePlayerIndex = i
		}
	}

	*game.activePlayer = game.players[(activePlayerIndex+1)%4]
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
