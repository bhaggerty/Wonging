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

func (d *Dealer) Initialize(id uint8, t *Table, s *Deck) {
	d.id = id
	d.table = t
	if s != nil {
		d.shoe = s
	} else {
		newShoe := new(Deck).Initialize(DEFAULTDECKPERSHOE)
		d.shoe = newShoe
	}
	d.shoe.Shuffle()
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

func (d *Dealer) PrintDealer() {
	fmt.Println("Dealer %d, sitting at table %d", d.id, d.table.id)
}

//Dealer actions
func (d *Dealer) dealSelf() {

}
func (d *Dealer) deal() {

}

//Dealer strategies
func (d *Dealer) standsOnAll17() {

}
func (d *Dealer) hitOnSoft17() {

}
