package value

import (
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
)

func CardValue(c *card.Card) int {
	switch c.Value {
	case card.Two:
		return 2
	case card.Three:
		return 3
	case card.Four:
		return 4
	case card.Five:
		return 5
	case card.Six:
		return 6
	case card.Seven:
		return 7
	case card.Eight:
		return 8
	case card.Nine:
		return 9
	case card.Ten, card.Jack, card.Queen, card.King:
		return 10
	default: // card.Ace
		return 11
	}
}
