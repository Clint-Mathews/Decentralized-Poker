package deck

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Hearts:
		return "HEARTS"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("Invalid card suit!")
	}
}

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	Suit  Suit
	Value int
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("Value of card cannot be higher than 13!")
	}
	return Card{
		Suit:  s,
		Value: v,
	}
}
