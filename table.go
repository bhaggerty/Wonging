package wonging

import (
// "fmt"
// "math/rand"
// "strconv"
// "time"
)

type Player struct {
	hand       *Hand
	currentBet float32
	totalCash  float32
}
type Dealer struct {
	shoe *Deck
}
type Table struct {
	//current count for all cards on table
	count int8

	//assuming one deck of card per dealer
	dealer *Dealer

	//Many players sitting at the table
	players []*Player
}

func (t *Table) GetTablePlayerNumber() int {
	return len(t.players)
}
