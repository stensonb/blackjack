package hand

import (
	"bytes"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
	"github.com/stensonb/blackjack/value"
)

type Hand struct {
	cards []*card.Card
}

func (h *Hand) String() string {

	var buffer bytes.Buffer

	for i := 0; i < len(h.cards); i++ {
		buffer.WriteString(h.cards[i].String())
		buffer.WriteString("\t")
	}

	return buffer.String()
}

func (h *Hand) Value() int {
	ans := 0
	aces := 0
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i].Value == card.Ace {
			aces += 1
		}
		ans += value.CardValue(h.cards[i])
	}

	// if we have aces, consider evaluating them as 1 instead of 11, if currently > 21 (busted)
	for ans > 21 {
		if aces > 0 {
			ans -= 10
			aces -= 1
			continue
		}
		break
	}

	return ans
}

// returns 0 if equal
// returns 1 if h > j
// returns -1 if h < j
func (h *Hand) Compare(j *Hand) int {
	hval := h.Value()
	jval := j.Value()

	switch {
	case hval == jval:
		return 0
	case hval > jval:
		return 1
	default: // hval < jval:
		return -1
	}
}

// add a card to this hand
func (h *Hand) AddCard(c *card.Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) Busted() bool {
	return h.Value() > 21
}

// returns the initial dealer hand, as shown before any players bet
// ie. one card showing, one card not showing
func (h *Hand) ShowOneCard() string {
	var buffer bytes.Buffer

	buffer.WriteString("XX\t")

	for i := 1; i < len(h.cards); i++ {
		buffer.WriteString(h.cards[i].String())
		buffer.WriteString("\t")
	}

	return buffer.String()
}
