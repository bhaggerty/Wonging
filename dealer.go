package wonging

type Dealer struct {
	id    uint8
	table *Table
	shoe  *Deck
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}
