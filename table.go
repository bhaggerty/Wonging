package wonging

import (
// "fmt"
// "math/rand"
// "strconv"
// "time"
)

type Player struct {
	id         uint8
	hand       *Hand
	currentBet float32
	totalCash  float32
}
type Dealer struct {
	id   uint8
	shoe *Deck
}
type Table struct {
	//current count for all cards on table
	count int8

	//assuming one deck of card per dealer
	dealer *Dealer

	//Many players sitting at the table
	players []*Player

	//limit on how many players can join
	playerLimit uint8
}

func (t *Table) GetTablePlayerNumber() int {
	return len(t.players)
}

func (t *Table) joinTable(p *Player) {
	if t.playerLimit > len(t.players) {
		t.players = append(t.players, p)
	} else {
		fmt.Println("Table is full, player " + p.id + " cannot join")
	}
}
