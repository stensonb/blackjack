package standarddeck

import (
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/deck"
	"math/rand"
	"testing"
)

const standardDeckSize = 52

func TestSize(t *testing.T) {
	d := New()
	if d.Size() != standardDeckSize { // standard deck size
		t.Error("standard deck has wrong number of cards")
	}

}

func TestNextCard(t *testing.T) {
	d := New()

	for i := 0; i < d.Size(); i++ {
		_, err := d.NextCard()
		if err != nil {
			t.Error("failed to get NextCard()")
		}
	}

	// now, we've used all the cards...so, we expect a NextCard() to error
	desiredError := deck.EmptyDeck{}

	_, err := d.NextCard()
	if err.Error() != desiredError.Error() {
		t.Error("wrong error (or nil) when the deck is empty")
	}

}

func TestCardsLeft(t *testing.T) {
	// test matrix of k,v pairs where k = # of calls to NextCard(), and v = expected CardsLeft()
	var matrix = map[int]int{0: standardDeckSize,
		int(standardDeckSize / 2): int(standardDeckSize / 2),
		standardDeckSize:          0,
		standardDeckSize + 1:      0}

	for k, v := range matrix {
		//get a new deck
		d := New()

		// call NextCard() k times
		for i := 0; i < k; i++ {
			d.NextCard()
		}

		//
		actual := d.CardsLeft()

		if actual != v {
			t.Error("wrong number of cards left.  should be", v, ", but received", actual)
		}
	}

}

// NOTE: this doesn't prove randomness, only that the order of the deck has been changed somehow
// additionally, there is a chance the shuffled deck is in the same order as the original, even
// after a successful shuffle -- just like in real life. :)
func TestShuffle(t *testing.T) {

	d0 := New()
	d1 := New()

	d0.Shuffle(rand.New(rand.NewSource(1))) // a static RNG seed for testing

	identical := true
	for i := 0; identical && i < d0.Size(); i++ {
		c0, _ := d0.NextCard()
		c1, _ := d1.NextCard()
		if !c0.Equals(c1) {
			identical = false
		}
	}

	if identical {
		t.Error("decks are identical, even after shuffling...")
	}
}
