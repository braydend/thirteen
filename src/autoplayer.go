package thirteen

import (
	"log"
	"math/rand"
	"sort"
)

func AutoPlay(hand []Card, format *Format, pile Pile) ([]Card, Format, error) {
	if *format == CLEAR {
		possibleFormats := []Format{}
		allPlays := buildPlaysForAllFormats(hand)
		for k, _ := range allPlays {
			possibleFormats = append(possibleFormats, k)
		}
		selectedFormat := possibleFormats[rand.Intn(len(possibleFormats))]

		return allPlays[selectedFormat], selectedFormat, nil
	}

	play, playedFormat := buildPlay(hand, *format, pile)

	if len(play) == 0 {
		return []Card{}, 0, nil
	}

	return play, playedFormat, nil
}

func buildPlaysForAllFormats(hand []Card) map[Format][]Card {
	plays := make(map[Format][]Card)
	emptyPile := NewPile()

	for i := 1; i < 25; i++ {
		format := Format(i)
		play, _ := buildPlay(hand, format, emptyPile)
		if len(play) > 0 {
			plays[format] = play
		}
	}

	return plays
}

func buildPlay(cards []Card, format Format, pile Pile) ([]Card, Format) {
	lastPlay := pile.LatestPlay().cards
	validPlays := [][]Card{}

	switch format {
	case SINGLE:
		fallthrough
	case PAIR:
		fallthrough
	case TRIPLE:
		fallthrough
	case QUAD:
		plays := buildMatchPlays(cards, int(format))
		for _, play := range plays {
			validPlays = append(validPlays, play)
		}
		break
	default:
		runLength := format - 2
		isFlush := runLength > 13
		if isFlush {
			runLength = runLength - 11
		}
		plays := buildRunPlays(cards, int(runLength), isFlush)
		for _, play := range plays {
			validPlays = append(validPlays, play)
		}
		break
	}

	sort.Slice(validPlays, func(i, j int) bool {
		iSorted := SortCards(validPlays[i])
		jSorted := SortCards(validPlays[j])

		iHighest := iSorted[len(iSorted)-1]
		jHighest := jSorted[len(jSorted)-1]

		if iHighest.Value == jHighest.Value {
			return iHighest.Suit < jHighest.Suit
		}

		return iHighest.Value < jHighest.Value
	})

	for _, play := range validPlays {
		if len(lastPlay) == 0 {
			return play, format
		}
		if isPlayHigherValue(lastPlay, play, format) {
			log.Printf("Playing: %v", StringifyCards(play))
			return play, format
		}
	}

	return []Card{}, 0
}

func isPlayHigherValue(previousPlay []Card, currentPlay []Card, format Format) bool {
	isFirstPlay := len(previousPlay) == 0 || format == CLEAR

	if isFirstPlay {
		return true
	}

	prevHighest := previousPlay[len(previousPlay)-1]
	curHighest := currentPlay[len(currentPlay)-1]

	if curHighest.Value < prevHighest.Value {
		return false
	}

	if curHighest.Value > prevHighest.Value {
		return true
	}

	if curHighest.Suit > prevHighest.Suit {
		return true
	}

	return false
}

func buildSinglePlays(hand []Card) [][]Card {
	sortedHand := SortCards(hand)
	plays := [][]Card{}

	for _, card := range sortedHand {
		plays = append(plays, []Card{card})
	}
	return plays
}

func buildMatchPlays(hand []Card, matchLength int) [][]Card {
	sortedHand := SortCards(hand)
	plays := [][]Card{}

	for i := range sortedHand {
		match := buildMatchFromOffset(hand, i, matchLength)

		if len(match) == matchLength {
			plays = append(plays, match)
		}
	}

	return plays
}

func buildRunPlays(hand []Card, runLength int, isFlush bool) [][]Card {
	sortedHand := SortCards(hand)
	plays := [][]Card{}

	for i := range sortedHand {
		run := buildRunFromOffset(hand, i, runLength, isFlush)

		if len(run) == runLength {
			plays = append(plays, run)
		}
	}

	return plays
}

func buildRunFromOffset(cards []Card, offset int, runLength int, isFlush bool) (out []Card) {
	offsetCards := cards[offset:]

	if isFlush {
		offsetCards = filterCards(cards, func(card Card) bool {
			return card.Suit == offsetCards[0].Suit
		})
	}

	if len(offsetCards) < runLength {
		return []Card{}
	}

	sortedCards := SortCards(offsetCards)

	sortedLength := len(sortedCards)

	for i, card := range sortedCards {
		if runLength == len(out) {
			return out
		}

		if i+1 <= sortedLength {
			if len(out) == 0 {
				out = append(out, card)
				continue
			}

			if out[len(out)-1].Value == card.Value {
				continue
			}

			if out[len(out)-1].Value+1 == card.Value {
				out = append(out, card)
			} else {
				return []Card{}
			}
		}
	}

	return out
}

func buildMatchFromOffset(cards []Card, offset int, matchLength int) (out []Card) {
	sortedCards := SortCards(cards[offset:])
	sameValueCards := []Card{}
	initialValue := sortedCards[0].Value
	for _, card := range sortedCards {
		if card.Value == initialValue {
			sameValueCards = append(sameValueCards, card)
		}
	}

	if len(sameValueCards) >= matchLength {
		return sameValueCards[:matchLength]
	}

	return out
}

func filterCards(cards []Card, test func(card Card) bool) []Card {
	filteredCards := []Card{}

	for _, card := range cards {
		if test(card) {
			filteredCards = append(filteredCards, card)
		}
	}

	return filteredCards
}
