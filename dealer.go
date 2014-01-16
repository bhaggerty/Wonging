package wonging

import (
	"fmt"
)

type Dealer struct {
	id                  uint8
	table               *Table
	shoe                *Deck
	curHand             *Hand
	faceDown            *Card
	action              DealerStrategy
	strategyDescription string
}

func (d *Dealer) Initialize(id uint8, t *Table, s *Deck) *Dealer {
	d.id = id
	d.table = t
	if s != nil {
		d.shoe = s
	} else {
		newShoe := new(Deck).Initialize(DEFAULTDECKPERSHOE)
		d.shoe = newShoe
	}
	d.curHand = new(Hand)
	d.action, d.strategyDescription = randomDealerStrategy()
	return d
}

func (d *Dealer) reset() {
	d.faceDown = nil
	d.curHand = new(Hand)
	d.action, d.strategyDescription = randomDealerStrategy()
}

// returns dealer's entire hand, combining facedown card with rest of hand
func (d *Dealer) fullHand() *Hand {
	if len(d.curHand.cards) != 0 && d.faceDown != nil {
		tmpHand := new(Hand)
		tmpHand.cards = append(d.curHand.cards, d.faceDown)
		return tmpHand
	}
	return nil
}

func (d *Dealer) changeTable(table *Table) {
	d.table = table
}

// returns total value of hand, combining facedown card and rest of hand
func (d *Dealer) calculateHandValue() (uint8, bool) {
	if len(d.curHand.cards) == 0 && d.faceDown == nil {
		//No card present, returning 0
		return 0, false
	}
	return d.fullHand().CalculateValue()
}

// returns value of hand, without facedown card
func (d *Dealer) calculateVisibleHandValue() (uint8, bool) {
	if len(d.curHand.cards) == 0 {
		return 0, false
	}
	return d.curHand.CalculateValue()
}

func (d *Dealer) isBusted() bool {
	tmpHand := new(Hand)
	tmpHand.cards = append(d.curHand.cards, d.faceDown)
	return tmpHand.isBusted()
}

//Dealer actions
func (d *Dealer) dealSelf() {
	if d.faceDown == nil {
		d.faceDown = d.shoe.pop()
	} else {
		d.curHand.cards = append(d.curHand.cards, d.shoe.pop())
	}
}

func (d *Dealer) deal() *Card {
	return d.shoe.pop()
}

func (d *Dealer) resetDeck() {
	d.shoe = new(Deck).Initialize(DEFAULTDECKPERSHOE)
	d.shoe.Shuffle()
}

// dealer simulation
func (d *Dealer) simulate() *Request {
	var req Request
	req.entityType = "dealer"
	req.id = d.id
	req.action = d.action(d)
	req.handIndex = []uint8{0}
	return &req
}
func (d *Dealer) PrintDealer() {
	fmt.Print(d.Description())
	if d.fullHand() != nil {
		d.fullHand().PrintHand()
	}
}

func (d *Dealer) Description() string {
	return fmt.Sprintf("[===== Dealer %d =====]\nstrategy: %s\n", d.id, CyanText(d.strategyDescription))
}

func (d *Dealer) GenerateHTMLMap() {
	html := "<html>"
	html += fmt.Sprintf("<h3>[===== Dealer %d =====]</h3>", d.id)
	html += fmt.Sprintf("<div>strategy: %s </div>", d.strategyDescription)
	html += "</html>"
	generateHTMLMap(fmt.Sprintf("dealer%d", d.id), html)
}
