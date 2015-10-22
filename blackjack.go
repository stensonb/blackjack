package main

import (
	"bytes"
	"fmt"
	"github.com/stensonb/playingcards/card"
	"github.com/stensonb/playingcards/standarddeck"
	"math/rand"
	"time"
)

type Hand struct {
	cards []*card.Card
}

func (h *Hand) String() string {

	var buffer bytes.Buffer

	for i := 0; i < len(h.cards); i++ {
		buffer.WriteString(h.cards[i].String())
		buffer.WriteString(" ")
	}

	return buffer.String()
}

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

func (h *Hand) Value() int {
	ans := 0
	for i := 0; i < len(h.cards); i++ {
		ans += CardValue(h.cards[i])
	}
	return ans
}

// returns 0 if equal
// returns 1 if h > h0
// returns -1 if h < h0
func (h *Hand) Compare(h0 *Hand) int {
	hval := h.Value()
	h0val := h0.Value()

	switch {
	case hval == h0val:
		return 0
	case hval > h0val:
		return 1
	default: // hval < h0val:
		return -1
	}
}

func main() {
	fmt.Println("Blackjack!")
	fmt.Println()

	d := standarddeck.New()
	d.Shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))

	dealer := new(Hand)
	player := new(Hand)

	c, _ := d.NextCard()
	player.cards = append(player.cards, c)
	c, _ = d.NextCard()
	dealer.cards = append(dealer.cards, c)
	c, _ = d.NextCard()
	player.cards = append(player.cards, c)
	c, _ = d.NextCard()
	dealer.cards = append(dealer.cards, c)

	fmt.Println("Dealer:", dealer)
	fmt.Println("Player:", player)

	fmt.Println() //formatting

	cmp := player.Compare(dealer)

	switch cmp {
	case 0:
		fmt.Println("Tie!")
	case 1:
		fmt.Println("Player wins!")
	case -1:
		fmt.Println("Dealer wins!")
	}

}
