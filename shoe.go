package wonging

import (
	"fmt"
	"math/rand"
	"strconv"
	// "time"
)

type Deck struct {
	cards []*Card
	dealt []*Card
}

func (d *Deck) Initialize(numOfDeck int) *Deck {
	suits := [4]string{"Diamonds", "Spades", "Hearts", "Clubs"} // unsorted
	for i := 0; i < numOfDeck; i++ {
		for _, suit := range suits {
			//take care of 2-10 first, their facevalues are the same as num value
			for i := 2; i <= 10; i++ {
				d.cards = append(d.cards, new(Card).Initialize(strconv.Itoa(i), uint8(i), suit))
			}
			//JQKA
			d.cards = append(d.cards, new(Card).Initialize("J", 10, suit))
			d.cards = append(d.cards, new(Card).Initialize("Q", 10, suit))
			d.cards = append(d.cards, new(Card).Initialize("K", 10, suit))
			d.cards = append(d.cards, new(Card).Initialize("A", 1, suit))
		}
	}
	d.Shuffle()
	return d
}

func (d *Deck) popRandom() *Deck {
	randomIndex := 0 + rand.Intn(len(d.cards))
	d.dealt = append(d.dealt, d.cards[randomIndex])
	d.cards = d.cards[:randomIndex+copy(d.cards[randomIndex:], d.cards[randomIndex+1:])]
	fmt.Printf("===== Deck: %d cards ===== Popped random: ", len(d.cards))
	d.cards[randomIndex].PrintCard()
	return d
}

func (d *Deck) pop() *Card {
	if len(d.cards) > 0 {
		cardPopped := d.cards[len(d.cards)-1]
		d.dealt = append(d.dealt, cardPopped)
		d.cards = d.cards[:len(d.cards)-1]
		fmt.Printf("===== Deck: %d cards ===== Popped from top: ", len(d.cards))
		cardPopped.PrintCard()
		return cardPopped
	}
	fmt.Println("Deck ran out of cards")
	return nil
}

func (d *Deck) Shuffle() {

	for i := 0; i < len(d.cards); i++ {
		r := rand.Intn(len(d.cards))
		temp := d.cards[i]
		d.cards[i] = d.cards[r]
		d.cards[r] = temp
	}
}

func (d *Deck) PrintDeck() {
	fmt.Println(d.Description())
	for _, card := range d.cards {
		card.PrintCard()
	}
}

func (d *Deck) Description() string {
	return fmt.Sprintf("===== Deck: %d cards =====", len(d.cards))
}
