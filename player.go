package wonging

import (
	"fmt"
)

type Player struct {
	id uint8
	//current table the player is sitting at
	table *Table
	//current casino
	casino *Casino
	//current hand, can be split into multiple hands
	hands []*Hand
	//how much is the player betting
	currentBet float64
	//if bought insurance for dealer getting blackjack
	isInsured []bool
	//if already doubled, index matching handIndex
	isDoubled []bool
	//if already surrendered
	isSurrendered []bool
	//how much money does the player have
	totalCash float64

	winCount  uint8
	loseCount uint8

	// play strategy
	action              PlayerStrategy
	strategyDescription string
	// count strategy
	count            CountingStrategy
	countDescription string

	// TODO: implement in phase 2, for group counting
	groupId uint8

	//TODO: implement in phase 2, for simulation of getting caught
	strikes uint8
}

func (p *Player) Initialize(id uint8, c *Casino, t *Table) *Player {
	p.id = id
	p.casino = c
	p.table = t
	p.currentBet = 0
	p.totalCash = DEFAULTPLAYERSTARTINGCASH
	//TODO: fix this fucking hack
	p.isInsured = []bool{false, false, false, false, false}
	p.isDoubled = []bool{false, false, false, false, false}
	p.isSurrendered = []bool{false, false, false, false, false}
	p.count, p.countDescription = randomCountingStrategy()
	if p.countDescription == "Ace/Five Count" {
		p.action = aceFiveCount
		p.strategyDescription = "Ace/Five Count"
	} else if p.countDescription == "Hi Lo" {
		p.action = hiLoCount
		p.strategyDescription = "Hi Lo Count"
	} else {
		p.action, p.strategyDescription = randomNonCountPlayerStrategy()
	}
	p.winCount = 0
	p.loseCount = 0
	return p
}

func (p *Player) reset() {
	//TODO: fix this fucking hack
	p.isInsured = []bool{false, false, false, false, false}
	p.isDoubled = []bool{false, false, false, false, false}
	p.isSurrendered = []bool{false, false, false, false, false}
	p.count, p.countDescription = randomCountingStrategy()
	if p.countDescription == "Ace/Five Count" {
		p.action = aceFiveCount
		p.strategyDescription = "Ace/Five Count"
	} else if p.countDescription == "Hi Lo" {
		p.action = hiLoCount
		p.strategyDescription = "Hi Lo Count"
	} else {
		p.action, p.strategyDescription = randomNonCountPlayerStrategy()
	}
	p.hands = nil
}

func (p *Player) bet(money float64, handIndex uint8) {
	if money <= 0 || p.totalCash < money {
		fmt.Println("No more money to make that bet")
		p.PrintPlayer()
	} else {
		p.hands[handIndex].bet(money)
		p.currentBet += money
		p.totalCash -= money
	}
}
func (p *Player) win(money float64) *Player {
	if money < 0 {
		fmt.Println("invalid amount, can't be negative!")
	} else {
		//get the current bet back
		p.totalCash += p.currentBet
		p.totalCash += money
		p.currentBet = 0
		p.casino.lose(money)
		p.winCount++
	}
	return p
}
func (p *Player) lose() *Player {
	p.casino.win(p.currentBet)
	p.currentBet = 0
	p.loseCount++
	return p
}

func (p *Player) profit() float64 {
	return p.totalCash - DEFAULTPLAYERSTARTINGCASH
}

func (p *Player) changeTable(table *Table) {
	p.table = table
}

func (p *Player) calculateHandValue(handIndex uint8) (uint8, bool) {
	value, soft := p.hands[handIndex].CalculateValue()
	return value, soft
}
func (p *Player) becomeObserver() bool {
	return p.table.playerBecomesObserver(p)
}
func (p *Player) leavesTable() bool {
	p.table = nil
	return p.casino.playerBecomesIdle(p)
}

func (p *Player) acceptCard(c *Card, handIndex uint8) {
	if handIndex > 0 {
		p.hands[handIndex].AddCard(c)
	} else {
		if p.hands == nil {
			p.hands = []*Hand{new(Hand)}
		}
		p.hands[0].AddCard(c)
	}
}

//player actions
// func (p *Player) hit(handIndex uint8) {
// 	if p.currentBet != 0 {
// 		p.table.playerRequest("hit", p, handIndex)
// 		fmt.Println("Player %d requesting a hit for hand: %d", p.id, handIndex)
// 	}
// }

func (p *Player) stand(handIndex uint8) {
	if p.currentBet != 0 {
		fmt.Println("Player %d is standing for hand: %d", p.id, handIndex)
	}
}

// func (p *Player) double(handIndex uint8) {
// 	if p.currentBet != 0 && !p.isDoubled {
// 		fmt.Println("Player %d is doubling his/her money, to %f, for hand: %d", p.id, p.currentBet, handIndex)
// 		p.bet(p.currentBet)
// 		p.hit(handIndex)
// 		p.isDoubled = true
// 	}
// }

func (p *Player) splitHand(handIndex uint8) {
	handToSplit := p.hands[handIndex]
	if p.currentBet != 0 && len(handToSplit.cards) == 2 && checkCardsValueEqual(handToSplit.cards[0], handToSplit.cards[1]) {
		//pointing new hand to second card
		var newHand *Hand
		newHand = new(Hand)
		newHand.AddCard(handToSplit.cards[1])
		p.hands = append(p.hands, newHand)
		//delete second card
		handToSplit.pop()
		//double the bet
		p.bet(p.hands[handIndex].currentBet, (uint8)(len(p.hands)-1))
	}
}
func (p *Player) splitAll() {
	if p.currentBet != 0 {
		for _, hand := range p.hands {
			if len(hand.cards) == 2 && checkCardsValueEqual(hand.cards[0], hand.cards[1]) {
				var newHand *Hand
				newHand = new(Hand)
				newHand.AddCard(hand.cards[1])
				p.hands = append(p.hands, newHand)
				//delete second card
				hand.pop()
			}
		}
	}
}

func (p *Player) surrenderAll() {
	save := p.currentBet / 2
	p.lose()
	p.win(save)
}

func (p *Player) surrender(handIndex uint8) {
	p.isSurrendered[handIndex] = true
}

func (p *Player) buyInsurance(handIndex uint8) {
	if p.currentBet != 0 && !p.isInsured[handIndex] {
		p.hands[handIndex].insure(p.hands[handIndex].currentBet / 2)
		p.isInsured[handIndex] = true
	} else {
		fmt.Println("Cannot buy insurance")
	}
}

// player state
func (p *Player) isBroke() bool {
	return p.currentBet+p.totalCash == 0
}

func (p *Player) isBanned() bool {
	return p.strikes > DEFAULTSTRIKEOUT
}

func (p *Player) isNatural() bool {
	return (len(p.hands) == 1 && p.hands[0].isBlackJack())
}

func (p *Player) isBusted(handIndex uint8) bool {
	return p.hands[handIndex].isBusted()
}
func (p *Player) isAllBusted() bool {
	for _, hand := range p.hands {
		if !(hand.isBusted()) {
			return false
		}
	}
	return true
}

// player simulation
func (p *Player) simulate() *Request {
	var req Request
	req.entityType = "player"
	req.id = p.id
	req.action, req.handIndex = p.action(p)
	return &req
}

func (p *Player) PrintPlayer() {

	fmt.Print(p.Description())
	if p.hands != nil && len(p.hands) > 0 {
		for _, hand := range p.hands {
			hand.PrintHand()
		}
	}
}

func (p *Player) Description() string {
	description := ""
	description += fmt.Sprintf("[===== Player %d =====]\n", p.id)
	description += fmt.Sprintf("currently betting: %f\n", p.currentBet)
	description += fmt.Sprintln("strategy: ", CyanText(p.strategyDescription))
	description += fmt.Sprintln("counting: ", CyanText(p.countDescription))

	if p.totalCash < DEFAULTPLAYERSTARTINGCASH {
		description += fmt.Sprintln("total cash: ", RedText(fmt.Sprintf("%f", p.totalCash)))
	} else {
		description += fmt.Sprintln("total cash: ", GreenText(fmt.Sprintf("%f", p.totalCash)))
	}

	if p.winCount < uint8(DEFAULTTOTALNUMBEROFGAMES/2) {
		description += fmt.Sprintln("winning: ", RedText(fmt.Sprintf("%d/%d", p.winCount, DEFAULTTOTALNUMBEROFGAMES)))
	} else {
		description += fmt.Sprintln("winning: ", GreenText(fmt.Sprintf("%d/%d", p.winCount, DEFAULTTOTALNUMBEROFGAMES)))
	}
	return description
}

func (p *Player) GenerateHTMLMap() {
	html := "<html>"
	html += fmt.Sprintf("<h3>[[===== Player %d =====]]</h3>", p.id)
	html += fmt.Sprintf("<div>currently betting: %f </div>", p.currentBet)
	html += fmt.Sprintf("<div>strategy: %s </div>", p.strategyDescription)
	html += fmt.Sprintf("<div>counting: %s </div>", p.countDescription)

	if p.totalCash < DEFAULTPLAYERSTARTINGCASH {
		html += fmt.Sprintln("total cash: ", fmt.Sprintf("<div><font color=\"red\">%f</font></div>", p.totalCash))
	} else {
		html += fmt.Sprintln("total cash: ", fmt.Sprintf("<div><font color=\"green\">%f</font></div>", p.totalCash))
	}

	if p.winCount < uint8(DEFAULTTOTALNUMBEROFGAMES/2) {
		html += fmt.Sprintln("winning: ", fmt.Sprintf("<div><font color=\"red\">%d/%d</font></div>", p.winCount, DEFAULTTOTALNUMBEROFGAMES))
	} else {
		html += fmt.Sprintln("winning: ", fmt.Sprintf("<div><font color=\"green\">%d/%d</font></div>", p.winCount, DEFAULTTOTALNUMBEROFGAMES))
	}

	html += "</html>"
	generateHTMLMap(fmt.Sprintf("player%d", p.id), html)
}
