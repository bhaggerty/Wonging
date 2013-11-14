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
	isInsured bool
	//if already doubled
	isDoubled bool
	//how much money does the player have
	totalCash float64

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
	p.isInsured = false
	p.isDoubled = false
	return p
}

func (p *Player) bet(money float64) {
	if money <= 0 || p.totalCash < money {
		fmt.Printf("No more money to make that bet")
		p.PrintPlayer()
	} else {
		p.currentBet += money
		p.totalCash -= money
	}
}
func (p *Player) win(money float64) {
	if money < 0 {
		fmt.Println("Use lose() instead!")
	} else {
		p.totalCash += money
		p.currentBet = 0
	}
}
func (p *Player) lose(money float64) {
	if p.totalCash >= money {
		p.totalCash -= money
	} else {
		p.totalCash = 0
	}
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
func (p *Player) hit(handIndex uint8) {
	if p.currentBet != 0 {
		p.table.playerRequest("hit", p, handIndex)
		fmt.Println("Player %d requesting a hit for hand: %d", p.id, handIndex)
	}
}

func (p *Player) stand(handIndex uint8) {
	if p.currentBet != 0 {
		fmt.Println("Player %d is standing for hand: %d", p.id, handIndex)
	}
}

func (p *Player) double(handIndex uint8) {
	if p.currentBet != 0 && !p.isDoubled {
		fmt.Println("Player %d is doubling his/her money, to %f, for hand: %d", p.id, p.currentBet, handIndex)
		p.bet(p.currentBet)
		p.hit(handIndex)
		p.isDoubled = true
	}
}

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

func (p *Player) buyInsurance() {
	if p.currentBet != 0 && !p.isInsured {
		p.bet(p.currentBet / 2)
		p.isInsured = true
	}
}

func (p *Player) isBroke() bool {
	return p.currentBet+p.totalCash == 0
}

func (p *Player) isBanned() bool {
	return p.strikes > DEFAULTSTRIKEOUT
}

func (p *Player) isNatural() bool {
	return (len(p.hands) == 1 && p.hands[0].isBlackJack())
}

func (p *Player) PrintPlayer() {
	fmt.Printf("[===== Player %d =====]\ncurrently betting: %f\ntotal cash: %f\n", p.id, p.currentBet, p.totalCash)
	if p.hands != nil && len(p.hands) > 0 {
		for _, hand := range p.hands {
			hand.PrintHand()
		}
	} else {
		fmt.Println("Player has no cards at the moment.")
	}
}
