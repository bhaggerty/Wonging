package wonging

import (
	"fmt"
)

type Casino struct {
	bank        float64
	tables      []*Table
	idleDealers []*Dealer
	idlePlayers []*Player
}

func (c *Casino) Initialize() *Casino {
	c.bank = DEFAULTCASINOSTARTINGCASH
	for i := 0; i < DEFAULTNUMBEROFTABLESPERCASINO; i++ {
		c.tables = append(c.tables, new(Table).Initialize(uint8(i)))
	}
	c.idleDealers = nil
	c.idlePlayers = nil
	return c
}

func (c *Casino) dealerBecomesIdle(d *Dealer) {
	dealerAlreadyIdle := false
	for _, iD := range c.idleDealers {
		if iD == d {
			dealerAlreadyIdle = true
		}
	}
	if dealerAlreadyIdle {
		fmt.Println("Dealer already idle")
	} else {
		c.idleDealers = append(c.idleDealers, d)
	}
}

func (c *Casino) playerBecomesIdle(d *Player) {
	playerAlreadyIdle := false
	for _, iP := range c.idlePlayers {
		if iP == d {
			playerAlreadyIdle = true
		}
	}
	if playerAlreadyIdle {
		fmt.Println("Player already idle")
	} else {
		c.idlePlayers = append(c.idlePlayers, d)
	}
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
