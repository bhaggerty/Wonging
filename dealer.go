package wonging

import (
	"fmt"
)

type Dealer struct {
	id    uint8
	table *Table
	shoe  *Deck
}

func (d *Dealer) Initialize(id uint8, t *Table, s *Deck) {
	d.id = id
	d.table = t
	if s != nil {
		d.shoe = s
	} else {
		newShoe := new(Deck).Initialize()
		d.shoe = newShoe
	}
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

func (d *Dealer) printDealer() {
	fmt.Println("Dealer %d, sitting at table %d", d.id, d.table.id)
}
