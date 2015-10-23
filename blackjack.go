package main

import (
	"fmt"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/deck"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/standarddeck"
	"github.com/stensonb/blackjack/hand"
	"math/rand"
	"time"
)

func NewBlackjack() (d *standarddeck.StandardDeck, dealer *hand.Hand, player *hand.Hand) {
	d = standarddeck.New()
	d.Shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))

	dealer = new(hand.Hand)
	player = new(hand.Hand)

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

func dealerPlays(d deck.Deck, h *hand.Hand) {
	for !h.Busted() && h.Value() < 17 { // dealer must hit when <17
		c, _ := d.NextCard()
		h.AddCard(c)
	}
}

func promptOptions(h *hand.Hand) string {
	// display options for player
	//TODO: make this more rich (if double-down or split is supported)
	return "[h]it, [s]tay: "
}

func promptPlayer(d deck.Deck, p *hand.Hand) bool {
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

}

func showHands(player, dealer *hand.Hand) {
	fmt.Println("Dealer:", dealer.ShowOneCard())
	fmt.Println("Player:", player)
	fmt.Println() //formatting
}

func declareWinner(player, dealer *hand.Hand) {
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
