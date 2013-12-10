package wonging

import (
	"fmt"
)

type Card struct {
	value       string
	numberValue uint8
	suit        string
	symbol      string
}

func (c *Card) Initialize(v string, n uint8, s string) *Card {
	c.value = v
	c.numberValue = n
	c.suit = s
	if s == "Diamonds" {
		c.symbol = RedText("♦")
	} else if s == "Hearts" {
		c.symbol = RedText("♥")
	} else if s == "Spades" {
		c.symbol = BlackText("♠")
	} else if s == "Clubs" {
		c.symbol = BlackText("♣")
	}
	return c
}

func (c *Card) PrintCard() {
	fmt.Println(c.value, c.symbol)
}
