package wonging

import (
	"fmt"
)

type Card struct {
	value       string
	numberValue uint8
	suit        string
}

func (c *Card) Initialize(v string, n uint8, s string) *Card {
	c.value = v
	c.numberValue = n
	c.suit = s
	return c
}

func (c *Card) PrintCard() {
	fmt.Println(c.value + " of " + c.suit)
}
