package wonging

import (
// "fmt"
// "math/rand"
// "strconv"
// "time"
)

const DEFAULTPLAYERLIMIT int8 = 5

type Table struct {
	id int8

	//current count for all cards on table
	count int8

	//assuming one deck of card per dealer
	dealer *Dealer

	//Many players sitting at the table
	players []*Player

	//limit on how many players can join
	playerLimit uint8
}

func (t *Table) Initialize(id int8) *Table {
	t.id = id
	t.count = 0
	t.dealer = nil
	t.players = nil
	t.playerLimit = DEFAULTPLAYERLIMIT
	return t
}

func (t *Table) GetTablePlayerNumber() int {
	return len(t.players)
}

func (t *Table) addPlayer(p *Player) {
	if t.playerLimit > len(t.players) {
		t.players = append(t.players, p)
	} else {
		fmt.Println("Table is full, player " + p.id + " cannot join")
	}
}
