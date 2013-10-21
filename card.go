package wonging

import (
	"fmt"
)

type Card struct {
	value       string
	numberValue int8
	suit        string
}

func (c *Card) NewCard(v string, n int8, s string) *Card {
	c.value = v
	c.numberValue = n
	c.suit = s
	return c
}

func (c *Card) printCard() {
	fmt.Println(c.value + " of " + c.suit)
}
