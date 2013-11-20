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
	action   DealerStrategy
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
	d.action = randomDealerStrategy()
	return d
}

func (d *Dealer) reset() {
	d.faceDown = nil
	d.curHand = new(Hand)
	d.action = randomDealerStrategy()
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

func (d *Dealer) calculateHandValue() (uint8, bool) {
	if len(d.curHand.cards) == 0 && d.faceDown == nil {
		//No card present, returning 0
		return 0, false
	}
	//combining cards into one hand
	tmpHand := new(Hand)
	tmpHand.cards = append(d.curHand.cards, d.faceDown)
	return tmpHand.CalculateValue()
}

func (d *Dealer) isBusted() bool {
	tmpHand := new(Hand)
	tmpHand.cards = append(d.curHand.cards, d.faceDown)
	return tmpHand.ifBusted()
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

func (d *Dealer) resetDeck() {
	d.shoe = new(Deck).Initialize(DEFAULTDECKPERSHOE)
	d.shoe.Shuffle()
}

// player simulation
func (d *Dealer) simulate() *Request {
	var req Request
	req.entityType = "dealer"
	req.id = d.id
	req.action = d.action(d)
	return &req
}
func (d *Dealer) PrintDealer() {
	fmt.Printf("[===== Dealer %d =====]\n", d.id)

	if d.faceDown != nil {
		value, soft := d.calculateHandValue()
		var softString string
		if soft {
			softString = "soft"
		} else {
			softString = "hard"
		}
		fmt.Printf("==> hand: (%s %d)\n", softString, value)
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
