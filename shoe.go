package wonging

import (
	"fmt"
)

type Card struct {
	value       string
	numberValue int8
	suit        string
}
type Deck struct {
	cards []*Card
}

func (c *Card) newCard(v string, n int8, s string) *Card {
	c.value = v
	c.numberValue = n
	c.suit = s
	return c
}

func (d *Deck) initDeck() *Deck {
	suits := []string{"Diamonds", "Spades", "Hearts", "Clubs"} // unsorted
	//take care of 2-10 first, their facevalues are the same as num value
	for _, suit := range suits {
		for i := 2; i <= 10; i++ {
			d.cards = append(d.cards, new(Card).newCard(string(i), int8(i), suit))
		}
	}
	fmt.Println(d)
	return d
}
