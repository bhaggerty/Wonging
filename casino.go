package wonging

import (
	"fmt"
)

type Casino struct {
	id          uint8
	bank        float64
	tables      []*Table
	idleDealers []*Dealer
	idlePlayers []*Player

	//TODO: implement in phase 2: for bannning card counters
	bannedPlayers []*Player
}

func (c *Casino) Initialize(id uint8) *Casino {
	c.id = id
	c.bank = DEFAULTCASINOSTARTINGCASH
	for i := 0; i < DEFAULTNUMBEROFTABLESPERCASINO; i++ {
		c.tables = append(c.tables, new(Table).Initialize(uint8(i), c))
	}
	for i := 0; i < DEFAULTNUMBEROFDEALERSPERCASINO; i++ {
		c.dealerBecomesIdle(new(Dealer).Initialize(uint8(i), nil, nil))
	}
	for i := 0; i < DEFAULTNUMBEROFPLAYERSPERCASINO; i++ {
		c.idlePlayers = append(c.idlePlayers, new(Player).Initialize(uint8(i), c, nil, nil))
	}
	return c
}

func (c *Casino) dealerBecomesIdle(d *Dealer) {
	if checkDealerContain(d, c.idleDealers) != -1 {
		fmt.Println("Dealer already idle")
	} else {
		c.idleDealers = append(c.idleDealers, d)
	}
}
func (c *Casino) dealerBecomesActive(d *Dealer) {
	if index := checkDealerContain(d, c.idleDealers); index == -1 {
		fmt.Println("Dealer not idling, cannot make him/her active")
	} else {
		c.idleDealers = append(c.idleDealers[:index], c.idleDealers[index+1:]...)

	}
}
func (c *Casino) playerBecomesIdle(p *Player) bool {
	if checkPlayerContain(p, c.idlePlayers) != -1 {
		fmt.Println("Player already idle")
		return false
	} else {
		c.idlePlayers = append(c.idlePlayers, p)
		return true
	}
}

func (c *Casino) playerBecomesActive(d *Player) bool {
	if index := checkPlayerContain(d, c.idlePlayers); index == -1 {
		fmt.Println("Player not idling, cannot make him/her active")
		return false
	} else {
		c.idlePlayers = append(c.idlePlayers[:index], c.idlePlayers[index+1:]...)
		return true
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

func (c *Casino) PrintCasino() {
	fmt.Printf("[[==== Casino %d ====]]\n", c.id)
	for _, table := range c.tables {
		table.printTable()
	}

}
