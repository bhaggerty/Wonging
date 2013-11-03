package main

import (
	"github.com/josephyzhou/wonging"
)

func main() {

	// deck := new(wonging.Deck).Initialize(1)
	// deck.PrintDeck()
	// deck = deck.pop()
	// deck = deck.DealRandom()
	// deck = deck.DealRandom()
	// deck = deck.Shuffle()
	casino := new(wonging.Casino).Initialize(0)
	casino.PrintCasino()
}
