package deck

import (
	"fmt"
	"math/rand"
	"strconv"
)

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

func (s Suit) SuitToUnicode() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
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

func (c Card) String() string {
	value := strconv.Itoa(c.Value)
	if value == "1" {
		value = "ACE"
	} else if value == "11" {
		value = "J"
	} else if value == "12" {
		value = "Q"
	} else if value == "13" {
		value = "K"
	}
	// Internally the call is requesting for a string so it calls `.String()`
	return fmt.Sprintf("%s of %s %s ", value, c.Suit, c.Suit.SuitToUnicode())
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

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x += 1
		}
	}
	return Shuffle(d)
}

func Shuffle(d Deck) Deck {
	for i := 0; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}
	return d
}
