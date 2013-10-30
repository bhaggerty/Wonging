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

	//game object to record results
	games []*Game
}

func (t *Table) Initialize(id uint8, c *Casino) *Table {
	t.id = id
	t.count = new(Counter).initialize()
	if c != nil {
		t.casino = c

	}
	t.playerLimit = DEFAULTPLAYERLIMITPERTABLE
	return t
}

func (t *Table) getNumberOfObservers() int {
	if t.idlePlayers != nil {
		return len(t.idlePlayers)
	} else {
		return 0
	}
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

func (t *Table) addPlayer(p *Player) bool {
	if t.playerLimit > uint8(len(t.players)) {
		t.players = append(t.players, p)
		p.changeTable(t)
		return true
	} else {
		fmt.Printf("Table is full, player %d cannot join", p.id)
		//putting player into observer array
		t.idlePlayers = append(t.idlePlayers, p)
		return false
	}
}
func (t *Table) addDealer(d *Dealer) bool {
	if t.dealer != nil {
		fmt.Printf("Table already has a dealer %d, he/she is now replaced by %d", t.dealer.id, d.id)
		//putting dealer into casino's idleDealer array
		t.casino.dealerBecomesIdle(t.dealer)
		return false
	}
	t.dealer = d
	t.dealer.changeTable(t)
	return true
}

func (t *Table) playerBecomesObserver(p *Player) bool {
	if checkPlayerContain(p, t.idlePlayers) != -1 {
		fmt.Println("Player is already observer")
		return false
	} else {
		t.idlePlayers = append(t.idlePlayers, p)
		return true
	}
}

func (t *Table) playerBecomesActive(p *Player) bool {
	if index := checkPlayerContain(p, t.idlePlayers); index == -1 {
		fmt.Println("Player not observing, cannot make him/her active")
		return false
	} else {
		t.idlePlayers = append(t.idlePlayers[:index], t.idlePlayers[index+1:]...)
		return true
	}
}

func (t *Table) calculateTableCount() *Counter {
	var allCounters []*Counter
	for _, player := range t.players {
		for _, hand := range player.hands {
			allCounters = append(allCounters, hand.calculateCount())
		}
	}
	allCounters = append(allCounters, t.dealer.curHand.calculateCount())
	return combineCounters(allCounters)
}

//pubsub
func (t *Table) playerRequest(action string, p *Player) {
	fmt.Println("Request from: %d action: %s", p.id, action)
	switch {
	case action == "hit":
		// p.acceptCard(t.dealer.deal())
	case action == "stand":
		//do stand
	}
}

func (t *Table) newGame() {
	fmt.Println("Table %d: Initializing a new game", t.id)
	game := new(Game).Initialize()
	t.games = append(t.games, game)
	for i := 1; i < 2; i++ {
		t.dealer.dealSelf()
		for _, player := range t.players {
			player.acceptCard(t.dealer.deal(), 0)
		}
	}
}

func (t *Table) printTable() {
	fmt.Printf("[===== Table %d =====]\n", t.id)
	t.dealer.PrintDealer()
	for _, player := range t.players {
		player.PrintPlayer()
	}

}
