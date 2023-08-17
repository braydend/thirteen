package thirteen

import (
	"fmt"
	"log"
)

type Pile struct {
	plays  []Play
	format Format
}

func NewPile() Pile {
	return Pile{
		plays:  []Play{},
		format: CLEAR,
	}
}

func (pile Pile) LatestPlay() Play {
	if len(pile.plays) == 0 {
		return NewPlay()
	}

	return pile.plays[len(pile.plays)-1]
}

func (pile *Pile) AddPlay(play Play) {
	pile.plays = append(pile.plays, play)
	log.Printf("%s", *pile)
}

func (pile Pile) Format() Format {
	return pile.format
}

func (pile *Pile) SetFormat(format Format) {
	pile.format = format
}

func (pile Pile) String() string {
	var out string

	out = fmt.Sprintf("%s--- START PILE ---\n", out)
	for i, play := range pile.plays {
		out = fmt.Sprintf("%s%d:%s\n", out, i+1, play)
	}
	out = fmt.Sprintf("%s--- END PILE ---\n", out)

	return out
}
