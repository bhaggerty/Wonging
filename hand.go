package wonging

const BLACKJACK int8 = 21

type Hand struct {
	cards []*Card
}

func (h *Hand) AddCard(c *Card) {
	h.cards = append(h.cards, c)
}
func (h *Hand) CalculateValue() int8 {
	var totalValue int8 = 0
	totalAs := 0
	if len(h.cards) == 0 {
		//No card present, returning 0
		return 0
	}
	for _, card := range h.cards {
		if card.value != "A" {
			totalValue += card.numberValue
		} else {
			//determine the optimal value of Aces later
			totalAs++
		}
	}
	for i := 0; i < totalAs; i++ {
		if totalValue > 10 {
			totalValue++
		} else {
			totalValue += 11
		}
	}
	return totalValue
}

func (h *Hand) CalculateCount() *Counter {
	counter := new(Counter).initialize()
	for _, card := range h.cards {
		counter = counter.count(card.value)
	}
	return counter
}

//optional parameter
//A total can be passed in, otherwise it will be calculated
func (h *Hand) ifBusted(total ...int8) bool {
	var myTotal int8
	if len(total) == 0 {
		myTotal = h.CalculateValue()
	} else if len(total) == 1 {
		myTotal = total[0]
	}
	if myTotal > BLACKJACK {
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
func (h *Hand) determineOutcome(totals ...int8) string {
	var myTotal int8
	if len(totals) == 2 {
		myTotal = totals[1]
	} else if len(totals) == 1 {
		myTotal = h.CalculateValue()
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
