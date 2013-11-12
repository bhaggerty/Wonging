package wonging

import (
	"fmt"
)

type Dealer struct {
	id       uint8
	table    *Table
	shoe     *Deck
	curHand  *Hand
	faceDown *Card
}

func (d *Dealer) Initialize(id uint8, t *Table, s *Deck) *Dealer {
	d.id = id
	d.table = t
	if s != nil {
		d.shoe = s
	} else {
		newShoe := new(Deck).Initialize(DEFAULTDECKPERSHOE)
		d.shoe = newShoe
	}
	d.curHand = new(Hand)
	d.shoe.Shuffle()
	return d
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

func (d *Dealer) calculateHandValue() uint8 {
	var totalValue uint8 = 0
	totalAs := 0
	if len(d.curHand.cards) == 0 && d.faceDown == nil {
		//No card present, returning 0
		return 0
	}
	//combining cards into one hand
	tmpHand := new(Hand)
	tmpHand.cards = append(d.curHand.cards, d.faceDown)
	for _, card := range tmpHand.cards {
		if card.value != "A" {
			totalValue += card.numberValue
		} else {
			//determine the optimal value of Aces later
			totalAs++
		}
	}
	for i := 0; i < totalAs; i++ {
		if totalValue > 10 || totalAs > 1 {
			totalValue++
		} else {
			totalValue += 11
		}
	}
	return totalValue
}

func (d *Dealer) PrintDealer() {
	fmt.Printf("[===== Dealer %d =====]\n", d.id)

	if d.faceDown != nil {
		fmt.Print("Facedown card: ")
		d.faceDown.PrintCard()
		if d.curHand != nil && d.curHand.cards != nil && len(d.curHand.cards) > 0 {
			for _, card := range d.curHand.cards {
				card.PrintCard()
			}
		}
	} else {
		fmt.Println("Dealer has no cards at the moment.")
	}

}

//Dealer actions
func (d *Dealer) dealSelf() {
	if d.faceDown == nil {
		d.faceDown = d.shoe.pop()
	} else {
		d.curHand.cards = append(d.curHand.cards, d.shoe.pop())
	}
}

// func (d *Dealer) deal(isDeal []bool) {
// 	for i := 0; i < len(d.table.players); i++ {
// 		if isDeal[i] {
// 			if d.table.players[i].currentBet != 0 {
// 				for _, hand := range d.table.players[i].hands {
// 					hand.AddCard(d.shoe.pop())
// 				}
// 			}
// 		}
// 	}
// }

func (d *Dealer) deal() *Card {
	return d.shoe.pop()
}
