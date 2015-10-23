package card

import (
	"testing"
)

func TestEquals(t *testing.T) {
	c := Card{Hearts, Ten}

	if !c.Equals(&Card{Hearts, Ten}) {
		t.Error("cards should equal")
	}

	if c.Equals(&Card{Spades, Ten}) {
		t.Error("cards should not equal")
	}

	if c.Equals(&Card{Spades, King}) {
		t.Error("cards should not equal")
	}
}
