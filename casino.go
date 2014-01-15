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
		c.idleDealers = append(c.idleDealers, new(Dealer).Initialize(uint8(i), nil, nil))
	}
	for i := 0; i < DEFAULTNUMBEROFPLAYERSPERCASINO; i++ {
		c.idlePlayers = append(c.idlePlayers, new(Player).Initialize(uint8(i), c, nil))
	}
	return c
}

func (c *Casino) DistributeDealers() *Casino {
	for i := 0; i < len(c.tables); i++ {
		c.tables[i].addDealer(c.idleDealers[0])
		c.idleDealers[0].changeTable(c.tables[i])
		c.dealerBecomesActive(c.idleDealers[0])
	}
	return c
}

func (c *Casino) DistributePlayers() *Casino {
	for i := 0; i < len(c.tables); i++ {
		for j := 0; j < int(c.tables[i].playerLimit); j++ {
			c.tables[i].addPlayer(c.idlePlayers[0])
			c.playerBecomesActive(c.idlePlayers[0])
		}
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

func (c *Casino) win(money float64) {
	c.bank += money
}
func (c *Casino) lose(money float64) {
	c.bank -= money
}

//calculate total cash flow
func (c *Casino) totalProfit() float64 {
	return (float64)(c.bank) - DEFAULTCASINOSTARTINGCASH
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
	fmt.Printf("\n[[======================================== PRINTING Casino %d ========================================]]\n", c.id)
	fmt.Print(c.Description())
	fmt.Println("[[===== Casino actives: =====]]")

	for _, table := range c.tables {
		table.PrintTable()
	}
	fmt.Println("[[==== Casino idles: ====]]")
	for _, idlePlayer := range c.idlePlayers {
		idlePlayer.PrintPlayer()
	}
	for _, idleDealer := range c.idleDealers {
		idleDealer.PrintDealer()
	}

}

func (c *Casino) Description() string {
	description := ""
	description += fmt.Sprintf("[[===== Casino %d =====]]\n", c.id)
	description += fmt.Sprintf("Total dealers: %d\n", c.totalDealers())
	description += fmt.Sprintf("Total players: %d\n", c.totalPlayers())
	description += fmt.Sprintf("Total idling dealers: %d\n", c.totalInactiveDealers())
	description += fmt.Sprintf("Total idling players: %d\n", c.totalInactivePlayers())
	description += fmt.Sprintf("Current cash: %f\n", c.bank)
	description += fmt.Sprintf("Current profit: %f\n", c.totalProfit())
	return description
}

func (c *Casino) Log() {
	logStr := "\n" + c.Description()
	for _, table := range c.tables {
		logStr += table.Description()
	}
	logStr += "[[==== Casino idles: ====]]\n"
	for _, idlePlayer := range c.idlePlayers {
		logStr += idlePlayer.Description()
	}
	for _, idleDealer := range c.idleDealers {
		logStr += idleDealer.Description()
	}
	logTo(fmt.Sprintf("Casino%d", c.id), logStr)
}

func (c *Casino) generateCasinoHTMLMap() string {
	html := "<html>"
	html += fmt.Sprintf("<h3>[[===== Casino %d =====]]</h3>", c.id)
	html += fmt.Sprintf("<h6>Total dealers: %d</h6>", c.totalDealers())
	html += fmt.Sprintf("<h6>Total players: %d</h6>", c.totalPlayers())
	html += fmt.Sprintf("<h6>Total idling dealers: %d</h6>", c.totalInactiveDealers())
	html += fmt.Sprintf("<h6>Total idling players: %d</h6>", c.totalInactivePlayers())
	html += fmt.Sprintf("<h6>Current cash: %f</h6>", c.bank)
	html += fmt.Sprintf("<h6>Current profit: %f</h6>", c.totalProfit())
	for _, table := range c.tables {
		html += fmt.Sprintf("<div><a href=\"table%d.html\">[Table %d]</a></div>", table.id, table.id)
	}
	html += "<h5>[[==== Casino idles: ====]]</h5>"
	for _, idlePlayer := range c.idlePlayers {
		html += fmt.Sprintf("<div><a href=\"player%d.html\">[Player %d]</a></div>", idlePlayer.id, idlePlayer.id)
	}
	for _, idleDealer := range c.idleDealers {
		html += fmt.Sprintf("<div><a href=\"dealer%d.html\">[Dealer %d]</a></div>", idleDealer.id, idleDealer.id)
	}
	html += "</html>"
	return html
}

//casino simulations
func (c *Casino) Start() {
	fmt.Println("Casino operating.")
	//do 200 games and see what happens
	for i := 0; i < DEFAULTTOTALNUMBEROFGAMES; i++ {
		for _, table := range c.tables {
			table.newGame(true)
		}
	}

}
