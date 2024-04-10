package main

import (
	"fmt"

	"github.com/Clint-Mathews/Decentralized-Poker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)

	fmt.Println("Hello Poker game!, card: ", card)
}
