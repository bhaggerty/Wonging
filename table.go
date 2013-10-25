package wonging

import (
	"fmt"

	// "math/rand"
	// "strconv"
// "time"
)

type Table struct {
	id uint8

	//current count for all cards on table
	count *Counter

	//which casino this table belongs to
	casino *Casino

	//assuming one deck of card per dealer
	dealer *Dealer

	//Many players sitting at the table
	players []*Player

	//observers
	idlePlayers []*Player

	//limit on how many players can join
	playerLimit uint8
}

func (t *Table) Initialize(id uint8, c *Casino) *Table {
	t.id = id
	t.count = new(Counter).initialize()
	if c != nil {
		t.casino = c

	} else {
		t.casino = nil
	}
	t.dealer = nil
	t.players = nil
	t.playerLimit = DEFAULTPLAYERLIMITPERTABLE
	return t
}

func (t *Table) getNumberOfPlayers() int {
	if t.players != nil {
		return len(t.players)
	} else {
		return 0
	}
}
func (t *Table) getNumberOfDealers() int {
	if t.dealer != nil {
		return 1
	} else {
		return 0
	}
}

func (t *Table) addPlayer(p *Player) {
	if t.playerLimit > uint8(len(t.players)) {
		t.players = append(t.players, p)
		p.changeTable(t)
	} else {
		fmt.Printf("Table is full, player %d cannot join", p.id)
	}
}
func (t *Table) addDealer(d *Dealer) {
	if t.dealer != nil {
		fmt.Printf("Table already has a dealer %d, he/she is now replaced by %d", t.dealer.id, d.id)
		t.dealer.changeTable(nil)
	}
	t.dealer = d
	t.dealer.changeTable(t)
}

func (t *Table) calculateTableCount() *Counter {
	var allCounters []*Counter
	for _, player := range t.players {
		allCounters = append(allCounters, player.hand.calculateCount())
	}
	return combineCounters(allCounters)
}
