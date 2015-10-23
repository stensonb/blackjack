package hand

import (
	"github.com/stensonb/blackjack/Godeps/_workspace/src/github.com/stensonb/playingcards/card"
	"testing"
)

func TestValue(t *testing.T) {
	h := new(Hand)
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.Ten})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Ten})
	if h.Value() != 20 {
		t.Error("value should equal 20")
	}

	h = new(Hand)
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.Jack})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Queen})
	if h.Value() != 20 {
		t.Error("value should equal 20")
	}

	h = new(Hand)
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.King})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Ace})
	if h.Value() != 21 {
		t.Error("value should equal 21")
	}

	h = new(Hand)
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.Four})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Five})
	if h.Value() != 9 {
		t.Error("value should equal 9")
	}
}

func TestCompare(t *testing.T) {
	//biggest
	biggest := new(Hand)
	biggest.AddCard(&card.Card{Suit: card.Spades, Value: card.Four})
	biggest.AddCard(&card.Card{Suit: card.Hearts, Value: card.Six})

	//smallest
	smallest := new(Hand)
	smallest.AddCard(&card.Card{Suit: card.Spades, Value: card.Three})
	smallest.AddCard(&card.Card{Suit: card.Hearts, Value: card.Five})

	if biggest.Compare(biggest) != 0 {
		t.Error("same hands should be equal")
	}

	if biggest.Compare(smallest) != 1 {
		t.Error("compare failed: biggest is bigger than smallest")
	}

	if smallest.Compare(biggest) != -1 {
		t.Error("compare failed: smallest is smaller than biggest")
	}
}

func TestAddCard(t *testing.T) {
	h := new(Hand)
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.Three})
	if h.Value() != 3 {
		t.Error("failed to addCard() to hand")
	}

}

func TestBusted(t *testing.T) {
	// 21
	h := new(Hand)
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Ten})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Ace})
	if h.Busted() {
		t.Error("should not have busted on 21")
	}

	// 22
	h = new(Hand)
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.Ten})
	h.AddCard(&card.Card{Suit: card.Hearts, Value: card.King})
	h.AddCard(&card.Card{Suit: card.Spades, Value: card.Two})
	if !h.Busted() {
		t.Error("should have busted on 22")
	}
}
