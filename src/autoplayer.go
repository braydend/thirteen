package thirteen

/*
requires:
- a hand
- reference to format
- reference to pile

will need to:
- check can play based on the format
- check play will beat latest play in pile
*/

type AutoPlayer struct {
	player *Player
	format *Format
	pile   *map[int][]Card
}

func NewAutoPlayer(player *Player, format *Format, pile *map[int][]Card) AutoPlayer {
	return AutoPlayer{player, format, pile}
}

func (ap *AutoPlayer) Play() bool {
	play := ap.BuildPlay()

	if len(play) == 0 {
		return false
	}

	return true
}

func (ap *AutoPlayer) BuildPlay() []Card {
	format := ap.format
	cards := ap.player.cards
	pileLength := len(*ap.pile)
	pile := *ap.pile
	lastPlay := pile[pileLength]

	validPlays := map[int][]Card{}

	switch *format {
	case SINGLE:
	case PAIR:
	case TRIPLE:
	case QUAD:
		plays := buildMatchPlays(*cards, int(*format))
		for i, play := range plays {
			validPlays[i] = play
		}
		break
	default:
		runLength := *format - 2
		isFlush := runLength > 13
		if isFlush {
			runLength = runLength - 13
		}
		plays := buildRunPlays(*cards, int(runLength), isFlush)
		for i, play := range plays {
			validPlays[i] = play
		}
		break
	}

	for _, play := range validPlays {
		if isPlayHigherValue(lastPlay, play) {
			return play
		}
	}

	return []Card{}
}

func isPlayHigherValue(previousPlay []Card, currentPlay []Card) bool {
	prevHighest := previousPlay[len(previousPlay)-1]
	curHighest := currentPlay[len(currentPlay)-1]

	if curHighest.Value > prevHighest.Value {
		return true
	}

	if curHighest.Suit > prevHighest.Suit {
		return true
	}

	return false
}

func buildSinglePlays(hand []Card) (plays map[int][]Card) {
	sortedHand := SortCards(hand)

	for i, card := range sortedHand {
		plays[i] = []Card{card}
	}
	return plays
}

func buildMatchPlays(hand []Card, matchLength int) (plays map[int][]Card) {
	sortedHand := SortCards(hand)

	for i := range sortedHand {
		match := buildMatchFromOffset(hand, i, matchLength)

		if len(match) == matchLength {
			index := len(plays)
			plays[index] = match
		}
	}

	return plays
}

func buildRunPlays(hand []Card, runLength int, isFlush bool) (plays map[int][]Card) {
	sortedHand := SortCards(hand)

	for i := range sortedHand {
		run := buildRunFromOffset(hand, i, runLength, isFlush)

		if len(run) == runLength {
			index := len(plays)
			plays[index] = run
		}
	}

	return plays
}

func buildRunFromOffset(cards []Card, offset int, runLength int, isFlush bool) (out []Card) {
	sortedCards := SortCards(cards[offset:])

	for i, card := range sortedCards {
		if runLength == len(out) {
			return out
		}
		if len(sortedCards) >= i+2 {
			isFlushCompatible := card.Suit == sortedCards[i+1].Suit

			if isNextCardValue(card, sortedCards[i+1]) {
				if isFlush && !isFlushCompatible {
					return []Card{}
				}
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
		return sameValueCards[:matchLength-1]
	}

	return out
}

func isNextCardValue(prev Card, next Card) bool {
	return prev.Value+1 == next.Value
}

func isSameCardValue(prev Card, next Card) bool {
	return prev.Value == next.Value
}
