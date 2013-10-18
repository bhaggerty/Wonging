package wonging

import (
	"fmt"
	"strconv"
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

func (c *Card) printCard() {
	fmt.Println(c.value + " of " + c.suit)
}

func (d *Deck) InitDeck() *Deck {
	suits := [4]string{"Diamonds", "Spades", "Hearts", "Clubs"} // unsorted
	for _, suit := range suits {
		//take care of 2-10 first, their facevalues are the same as num value
		for i := 2; i <= 10; i++ {
			d.cards = append(d.cards, new(Card).newCard(strconv.Itoa(i), int8(i), suit))
		}
		//JQKA
		d.cards = append(d.cards, new(Card).newCard("J", 10, suit))
		d.cards = append(d.cards, new(Card).newCard("Q", 10, suit))
		d.cards = append(d.cards, new(Card).newCard("K", 10, suit))
		d.cards = append(d.cards, new(Card).newCard("A", 1, suit))
	}
	d.PrintDeck()
	return d
}

func (d *Deck) PrintDeck() {
	fmt.Println("There are " + strconv.Itoa(len(d.cards)) + " cards in the deck\n===============================\n")
	for _, card := range d.cards {
		card.printCard()
	}
}
