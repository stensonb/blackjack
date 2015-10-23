package card

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// note the distict absence of numerical value to these card "values"
// this is intentionally detached so that any given game can assign "rank"
// to a card's value.
const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Suit int

func (s *Suit) String() string {
	switch *s {
	case Spades:
		return "♠"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	default: // Clubs
		return "♣"
	}
}

type Value int

func (v *Value) String() string {
	switch *v {
	case Ace:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	default: // King
		return "K"
	}
}

var AllSuits = []Suit{Hearts, Spades, Diamonds, Clubs}
var AllValues = []Value{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type Card struct {
	Suit  Suit
	Value Value
}

func (c *Card) String() string {
	return c.Value.String() + c.Suit.String()
}

func (c *Card) Equals(o *Card) bool {
	if c.Suit == o.Suit && c.Value == o.Value {
		return true
	}
	return false
}
