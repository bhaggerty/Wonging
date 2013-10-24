package wonging

import (
	"fmt"
)

type Player struct {
	id uint8
	//current table the player is sitting at
	table *Table
	//current hand
	hand *Hand
	//how much is the player betting
	currentBet float32
	//how much money does the player have
	totalCash float32

	//TODO: implement in phase 2, for group counting
	//groupId uint8

	//TODO: implement in phase 2, for simulation of getting caught
	//strike uint8
}

func (p *Player) bet(money float32) {
	if money <= 0 || (p.totalCash-money) < 0 {
		fmt.Println("Invalid bet")
	} else {
		p.currentBet += money
		p.totalCash -= money
	}
}

func (p *Player) changeTable(table *Table) {
	p.table = table
}

func (p *Player) acceptCard(c *Card) {
	p.hand.AddCard(c)
}

func (p *Player) printPlayer() {
	fmt.Println("Player %d, sitting at table %d, currently betting %f, total cash: %f", p.id, p.table.id, p.currentBet, p.totalCash)
}
