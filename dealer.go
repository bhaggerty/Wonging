package wonging

import (
	"fmt"
)

type Dealer struct {
	id    uint8
	table *Table
	shoe  *Deck
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

func (d *Dealer) printDealer() {
	fmt.Println("Dealer %d, sitting at table %d", d.id, d.table.id)
}
