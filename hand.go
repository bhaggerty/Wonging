package wonging

type Hand struct {
	cards []*Card
}

func (h *Hand) AddCard(c *Card) {
	h.cards = append(h.cards, c)
}
func (h *Hand) pop() {
	h.cards = h.cards[:len(h.cards)-1]
}

//calculate the value of current hand
//returns:
//	totalValue: uint8 [value of hand]
//	soft: bool [if the hand is considered soft]
func (h *Hand) CalculateValue() (uint8, bool) {
	var totalValue uint8 = 0
	var soft bool = false
	totalAs := 0
	if len(h.cards) == 0 {
		//No card present, returning 0
		return 0, false
	}
	for _, card := range h.cards {
		if card.value != "A" {
			totalValue += card.numberValue
		} else {
			//determine the optimal value of Aces later
			totalAs++
		}
	}
	if totalAs > 0 {
		if totalValue > 10 {
			totalValue += uint8(totalAs)
		} else {
			if totalValue+11+uint8(totalAs-1) <= BLACKJACK {
				totalValue = totalValue + 11 + uint8(totalAs-1)
				soft = true
			} else {
				totalValue += uint8(totalAs)
			}
		}
	}
	return totalValue, soft
}
func (h *Hand) isBlackJack() bool {
	if len(h.cards) == 2 {
		return uint8(h.cards[0].numberValue)+uint8(h.cards[1].numberValue) == BLACKJACK
	} else {
		return false
	}

}

func (h *Hand) calculateCount() *Counter {
	counter := new(Counter).initialize()
	for _, card := range h.cards {
		counter = counter.count(card.value)
	}
	return counter
}

//optional parameter
//A total can be passed in, otherwise it will be calculated
func (h *Hand) ifBusted(total ...uint8) bool {
	var myTotal uint8
	if len(total) == 0 {
		myTotal, _ = h.CalculateValue()
	} else if len(total) == 1 {
		myTotal = total[0]
	}
	if uint8(myTotal) > BLACKJACK {
		return true
	} else {
		return false
	}
}

//optional parameter
//One or two total(s) can be passed in
//two totals passed in: we will do the comparison right away
//One total passed in: assuming it is the total of the hand of opponent
//                     proceed to calculating own total then compare
func (h *Hand) DetermineOutcome(totals ...uint8) string {
	var myTotal uint8
	if len(totals) == 2 {
		myTotal = totals[1]
	} else if len(totals) == 1 {
		myTotal, _ = h.CalculateValue()
	} else {
		return "Pass in at least one, but not more than two totals for comparison"
	}
	dealerTotal := totals[0]
	if h.ifBusted(myTotal) {
		return "Player busted"
	} else if h.ifBusted(dealerTotal) {
		return "Dealer busted"
	} else if myTotal == dealerTotal {
		return "Push"
	} else if myTotal > dealerTotal {
		return "Player wins"
	} else {
		return "Dealer wins"
	}
}
