package wonging

type Casino struct {
	tables      []*Table
	idleDealers []*Dealer
	idlePlayers []*Player
}

//calculate total cash flow
func (c *Casino) totalCashFlow() float32 {
	//TODO
	return 0.1
}

//calculate total profit
func (c *Casino) totalProfit() float32 {
	//TODO
	return 0.1
}

//total dealers
func (c *Casino) totalDealers() uint8 {
	return c.totalActiveDealers() + c.totalInactiveDealers()
}

func (c *Casino) totalActiveDealers() uint8 {
	total := 0
	for _, table := range c.tables {
		total += table.getNumberOfDealers()
	}
	return uint8(total)
}

func (c *Casino) totalInactiveDealers() uint8 {
	return uint8(len(c.idleDealers))
}

//total players
func (c *Casino) totalPlayers() uint8 {
	return c.totalActivePlayers() + c.totalInactivePlayers()
}

func (c *Casino) totalActivePlayers() uint8 {
	total := 0
	for _, table := range c.tables {
		total += table.getNumberOfPlayers()
	}
	return uint8(total)
}

func (c *Casino) totalInactivePlayers() uint8 {
	return uint8(len(c.idlePlayers))
}
