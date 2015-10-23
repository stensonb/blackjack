package main

import (
	"bytes"
	"fmt"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/deck"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/standarddeck"
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
		buffer.WriteString("\t")
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
	aces := 0
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i].Value == card.Ace {
			aces += 1
		}
		ans += CardValue(h.cards[i])
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
func showOneCard(h *Hand) string {
	var buffer bytes.Buffer

	buffer.WriteString("XX\t")

	for i := 1; i < len(h.cards); i++ {
		buffer.WriteString(h.cards[i].String())
		buffer.WriteString("\t")
	}

	return buffer.String()
}

func NewBlackjack() (d *standarddeck.StandardDeck, dealer *Hand, player *Hand) {
	d = standarddeck.New()
	d.Shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))

	dealer = new(Hand)
	player = new(Hand)

	c, _ := d.NextCard()
	player.AddCard(c)
	c, _ = d.NextCard()
	dealer.AddCard(c)
	c, _ = d.NextCard()
	player.AddCard(c)
	c, _ = d.NextCard()
	dealer.AddCard(c)

	return d, dealer, player

}

func dealerPlays(d deck.Deck, h *Hand) {
	for !h.Busted() && h.Value() < 17 { // dealer must hit when <17
		c, _ := d.NextCard()
		h.AddCard(c)
	}
}

func promptOptions(h *Hand) string {
	// display options for player
	//TODO: make this more rich (if double-down or split is supported)
	return "[h]it, [s]tay: "
}

func promptPlayer(d deck.Deck, p *Hand) bool {
	fmt.Print(promptOptions(p))

	// read from the player
	switch getUserInput() {
	case "h":
		c, _ := d.NextCard()
		p.AddCard(c)
		return false
	default: // case "s":
		return true
	}

	//	answer := getUserInput()
	//	fmt.Println(answer)

}

func getUserInput() string {
	var input string

	for {
		fmt.Scanln(&input)
		fmt.Println() //formatting
		switch input {
		case "h", "s":
			return input
		default:
			fmt.Println("try again")
		}
	}

	fmt.Scanln(&input)
	return input
}

func showHands(player, dealer *Hand) {
	fmt.Println("Dealer:", showOneCard(dealer))
	fmt.Println("Player:", player)
	fmt.Println() //formatting
}

func declareWinner(player, dealer *Hand) {
	fmt.Println("Dealer:", dealer)
	fmt.Println("Player:", player)
	fmt.Println() //formatting

	cmp := player.Compare(dealer)

	switch {
	case player.Busted() && dealer.Busted(): // house rules
		fmt.Println("Dealer wins!")
	case player.Busted():
		fmt.Println("Dealer wins!")
	case dealer.Busted():
		fmt.Println("Player wins!")
	case cmp == 0:
		fmt.Println("Tie!")
	case cmp == 1:
		fmt.Println("Player wins!")
	case cmp == -1:
		fmt.Println("Dealer wins!")
	}
}

func main() {
	fmt.Println("----------")
	fmt.Println("Blackjack!")
	fmt.Println() // formatting

	d, dealer, player := NewBlackjack()

	// player loop
	for done := false; done == false && !player.Busted(); done = promptPlayer(d, player) {
		showHands(player, dealer)
	}

	// dealer loop
	dealerPlays(d, dealer)

	declareWinner(player, dealer)
	fmt.Println()

}
