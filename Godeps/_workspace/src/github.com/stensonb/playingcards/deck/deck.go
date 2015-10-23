package deck

import (
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
	"math/rand"
)

// an interface to define common behaviors of a Deck
// look at github.com/stensonb/playingcards/standarddeck
// for a sample implementation which you can use
type Deck interface {
	Shuffle(*rand.Rand)            // shuffle the deck
	NextCard() (*card.Card, error) // issue the next card from the top
	Size() int                     // total number of cards in this deck started with
	CardsLeft() int                // number of remaining cards in the deck
}

// an error used to signify the deck has no more cards
type EmptyDeck struct{}

// satisfing the error interface
func (e EmptyDeck) Error() string {
	return "Deck is empty.  No more cards."
}
