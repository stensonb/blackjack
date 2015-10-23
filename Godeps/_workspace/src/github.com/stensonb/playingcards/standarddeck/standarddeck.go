package standarddeck

import (
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/deck"
	"math/rand"
)

type StandardDeck struct {
	cards []*card.Card
	idx   int
}

// returns a standard deck, all values of all four suits, unshuffled.
func New() *StandardDeck {
	ans := new(StandardDeck)

	// add all cards from all suits
	for _, sv := range card.AllSuits {
		for _, vv := range card.AllValues {
			c := card.Card{Suit: sv, Value: vv}
			ans.cards = append(ans.cards, &c)
		}
	}

	return ans
}

// iterate through the deck, and swap values with
// a random card from another location in the deck
func (sd *StandardDeck) Shuffle(r *rand.Rand) {
	s := sd.Size()
	for k := range sd.cards {
		i := r.Intn(s)
		sd.cards[k], sd.cards[i] = sd.cards[i], sd.cards[k]
	}
}

// returns the next card from the "top" of the deck
// if there are none left, return deck.EmptyDeck as the error
func (sd *StandardDeck) NextCard() (*card.Card, error) {
	var ans *card.Card

	if sd.idx < sd.Size() { // ensure we're not at the end of the cards array
		ans = sd.cards[sd.idx]
		sd.idx += 1 // go to the next card
		return ans, nil
	}

	// otherwise
	return &card.Card{}, deck.EmptyDeck{}
}

// the size, as an integer, of the number of cards in this deck
func (sd *StandardDeck) Size() int {
	return len(sd.cards)
}

// number of remaining cards in the deck
func (sd *StandardDeck) CardsLeft() int {
	return sd.Size() - sd.idx
}
