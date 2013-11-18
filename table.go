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
	//get visible cards from dealer as well
	allCounters = append(allCounters, t.dealer.curHand.calculateCount())
	return combineCounters(allCounters)
}

//pubsub
// func (t *Table) playerRequest(action string, p *Player, handIndex uint8) {
// 	fmt.Println("Request from: %d action: %s", p.id, action)
// 	switch {
// 	case action == "hit":
// 		p.acceptCard(t.dealer.deal(), handIndex)
// 	case action == "stand":
// 		//do stand
// 	}
// }

func (t *Table) newGame(resetDeck bool) {
	fmt.Printf("Table %d: Initializing a new game.\n", t.id)

	//player betting amounts
	//and reset other settings
	for _, player := range t.players {
		player.bet(DEFAULTPLAYERBET)
		player.reset()
	}
	// take care of resetDeck
	if resetDeck {
		t.dealer.resetDeck()
	}
	// dealer reset his/her cards
	t.dealer.reset()

	//deal cards to dealers and players, 2/person
	for i := 0; i < 2; i++ {
		t.dealer.dealSelf()
		for _, player := range t.players {
			player.acceptCard(t.dealer.deal(), 0)
		}
	}
	game := new(Game).Initialize(t)
	t.games = append(t.games, game)
}

//main engine of the entire project
//TODO: Might need to find a new place to put this
func (t *Table) simulate() {
	doneCount := 0
	for doneCount < (t.getNumberOfPlayers() + t.getNumberOfDealers()) {
		doneCount = 0
		playerRequestQueue := make(chan *Request, len(t.players))
		//players simulations - order matters here, no go routine
		for i := 0; i < len(t.players); i++ {
			playerRequestQueue <- t.players[i].simulate()
			select {
			case req := <-playerRequestQueue:
				req.printRequest()
				switch req.action {
				case "stand":
					doneCount++
				case "hit":
					if t.players[i].currentBet != 0 {
						t.players[i].acceptCard(t.dealer.deal(), req.handIndex)
					}
				case "double":
					if t.players[i].currentBet != 0 && !t.players[i].isDoubled {
						//bet same money
						t.players[i].bet(t.players[i].currentBet)
						t.players[i].isDoubled = true

						//hit
						t.players[i].acceptCard(t.dealer.deal(), req.handIndex)
					}
				case "splitHand":

				case "splitAllHands":

				}
			}
		}
		close(playerRequestQueue)
		dealerRequestQueue := make(chan *Request, 1)
		dealerRequestQueue <- t.dealer.simulate()
		select {
		case req := <-dealerRequestQueue:
			req.printRequest()
			switch req.action {
			case "stand":
				doneCount++
			case "dealSelf":
				t.dealer.dealSelf()
			}
		}

		/** update the game object, end of round **/

		//get last game obj = current
		curGame := t.games[len(t.games)-1]
		curGame.round++
	}
	t.determineOutcome()
	t.games[len(t.games)-1].updatePlayerResult(t)
}

func (t *Table) determineOutcome() {
	fmt.Println("==[[ Trying to determine outcome of this game ]]==")
	// case player busted
	var remainingPlayers = []*Player{}
	for _, player := range t.players {
		if player.isAllBusted() {
			fmt.Printf("Player %d Busted\n", player.id)
			player.lose()
		} else {
			remainingPlayers = append(remainingPlayers, player)
		}
	}
	if len(remainingPlayers) > 0 {
		// case dealer is busted, everyone not busted get 3/2 payout
		if t.dealer.isBusted() {
			fmt.Printf("Dealer %d Busted\n", t.dealer.id)
			for _, player := range remainingPlayers {
				player.win(player.currentBet * 0.5)
			}
		} else {
			for _, player := range remainingPlayers {
				for i := 0; i < len(player.hands); i++ {
					pValue, _ := player.calculateHandValue(uint8(i))
					dValue, _ := t.dealer.calculateHandValue()
					if pValue == dValue {
						fmt.Printf("Player %d same value, push? [%d vs %d]\n", player.id, pValue, dValue)
					} else if pValue > dValue {
						fmt.Printf("Player %d beats dealer hand [%d vs %d]\n", player.id, pValue, dValue)
						player.win(player.currentBet * 0.5)
					} else {
						fmt.Printf("Player %d loses to dealer's hand [%d vs %d]\n", player.id, pValue, dValue)
						player.lose()
					}
				}
			}
		}
	}
}

func (t *Table) PrintTable() {
	fmt.Printf("[===== Table %d =====]\n", t.id)
	fmt.Printf("[Number of dealers: %d]\n", t.getNumberOfDealers())
	fmt.Printf("[Number of players: %d]\n", t.getNumberOfPlayers())
	fmt.Printf("[Number of observers: %d]\n", t.getNumberOfObservers())
	if t.getNumberOfDealers() > 0 {
		t.dealer.PrintDealer()
	}
	if t.getNumberOfPlayers() > 0 {
		for _, player := range t.players {
			player.PrintPlayer()
		}
	}
	if t.games != nil && len(t.games) > 0 {
		t.games[len(t.games)-1].PrintGame()
	}
}
