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
	switch s {
	case "Diamonds":
		c.symbol = RedText("♦")
	case "Hearts":
		c.symbol = RedText("♥")
	case "Spades":
		c.symbol = BlackText("♠")
	case "Clubs":
		c.symbol = BlackText("♣")
	}
	return c
}

func (c *Card) PrintCard() {
	fmt.Println(c.Description())
}

func (c *Card) Description() string {
	return fmt.Sprint(c.value, c.symbol)
}
