package main

import (
	"github.com/josephyzhou/wonging"
)

func main() {

	deck := new(wonging.Deck).Initialize(1)
	deck.PrintDeck()
	// deck = deck.pop()
	// deck = deck.DealRandom()
	// deck = deck.DealRandom()
	// deck = deck.Shuffle()
}
