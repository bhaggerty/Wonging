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
func (t *Table) calculateTotalNumberOfHands() int {
	sum := 0
	for _, player := range t.players {
		sum += len(player.hands)
	}
	return sum
}

func (t *Table) newGame(resetDeck bool) {
	fmt.Println(WhiteText(fmt.Sprintf("\n==[[ Table %d: Initializing a new game. ]]==", t.id)))

	//player betting amounts
	//and reset other settings
	for _, player := range t.players {
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

	//place bet on hand
	for _, player := range t.players {
		player.bet(DEFAULTPLAYERBET, 0)
	}
	game := new(Game).Initialize(t)
	t.games = append(t.games, game)

	// start simulation
	t.simulate()
}

//main engine of the entire project
//TODO: Might need to find a new place to put this
func (t *Table) simulate() {
	doneCount := 0
	dealerDone := false
	for doneCount < t.calculateTotalNumberOfHands() {
		playerRequestQueue := make(chan *Request, len(t.players))
		doneCount = 0
		//players simulations - order matters here, no go routine
		for i := 0; i < len(t.players); i++ {
			playerRequestQueue <- t.players[i].simulate()
			select {
			case req := <-playerRequestQueue:
				req.printRequest()
				if t.players[i].currentBet > 0 {
					currentPlayer := t.players[i]
					for j := 0; j < len(req.action); j++ {

						switch req.action[j] {
						case "stand":
							doneCount++
						case "surrender":
							doneCount++
							currentPlayer.surrender(req.handIndex[j])
						case "hit":
							currentPlayer.acceptCard(t.dealer.deal(), req.handIndex[j])
						case "double":
							if !currentPlayer.isDoubled[j] {
								//bet same money
								currentPlayer.bet(currentPlayer.hands[req.handIndex[j]].currentBet, req.handIndex[j])
								currentPlayer.isDoubled[j] = true
							}
							//hit
							currentPlayer.acceptCard(t.dealer.deal(), req.handIndex[j])
						case "split":
							currentPlayer.splitHand(req.handIndex[j])
							currentPlayer.acceptCard(t.dealer.deal(), req.handIndex[j])
							currentPlayer.acceptCard(t.dealer.deal(), uint8(len(currentPlayer.hands)-1))
						case "splitAllHands":
							currentPlayer.splitAll()
						default:
							fmt.Println("invalid action")
							doneCount++
						}
					}

				}
			}
		}
		fmt.Println("Done count, calculateTotalNumberOfHands: ", doneCount, " ", t.calculateTotalNumberOfHands())

		close(playerRequestQueue)
	}
	for !dealerDone {
		dealerRequestQueue := make(chan *Request, 1)
		dealerRequestQueue <- t.dealer.simulate()
		select {
		case req := <-dealerRequestQueue:
			req.printRequest()
			for _, action := range req.action {
				switch action {
				case "stand":
					dealerDone = true
				case "dealSelf":
					t.dealer.dealSelf()
				}
			}

		}
	}
	t.determineOutcome()
	t.games[len(t.games)-1].updatePlayerResult(t)
}

func (t *Table) determineOutcome() {
	fmt.Println(WhiteText("==[[ Trying to determine outcome of this game ]]=="))
	// case player busted
	var remainingPlayers = []*Player{}
	for _, player := range t.players {
		if player.isAllBusted() {
			fmt.Println(RedText(fmt.Sprintf("Player %d Busted", player.id)))
			player.lose()
		} else {
			remainingPlayers = append(remainingPlayers, player)
		}
	}
	if len(remainingPlayers) > 0 {
		for _, player := range remainingPlayers {
			// case player is natural blackjack, dealer is not, 3:2 payout
			if player.isNatural() && !t.dealer.curHand.isBlackJack() {
				player.win(player.currentBet * 1.5)
			}
			// handle player surrender hand(s)
			for i := 0; i < len(player.hands); i++ {
				if player.isSurrendered[i] {
					save := player.hands[i].currentBet / 2
					player.lose()
					player.win(save)
				}
			}
		}

		// case dealer is busted, everyone not busted get 1:1 payout
		if t.dealer.isBusted() {
			fmt.Println(GreenText(fmt.Sprintf("Dealer %d Busted", t.dealer.id)))
			for _, player := range remainingPlayers {
				player.win(player.currentBet)
			}
		} else {
			for _, player := range remainingPlayers {
				for i := 0; i < len(player.hands); i++ {
					pValue, _ := player.calculateHandValue(uint8(i))
					dValue, _ := t.dealer.calculateHandValue()
					if pValue == dValue {
						fmt.Println(RedText(fmt.Sprintf("Player %d same value, dealer wins [%d vs %d]", player.id, pValue, dValue)))
						player.lose()
					} else if pValue > dValue {
						fmt.Println(GreenText(fmt.Sprintf("Player %d beats dealer hand [%d vs %d]", player.id, pValue, dValue)))
						player.win(player.currentBet)
					} else {
						fmt.Println(RedText(fmt.Sprintf("Player %d loses to dealer's hand [%d vs %d]", player.id, pValue, dValue)))
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

func (t *Table) Description() string {
	description := ""
	description += fmt.Sprintf("[===== Table %d =====]\n", t.id)
	description += fmt.Sprintf("[Number of dealers: %d]\n", t.getNumberOfDealers())
	description += fmt.Sprintf("[Number of players: %d]\n", t.getNumberOfPlayers())
	description += fmt.Sprintf("[Number of observers: %d]\n", t.getNumberOfObservers())
	return description
}
