//similar to dealer strategies class
//To use: first init a Player's strategy in init function
//call simulate(), which would in turn call one of these strategy functions
package wonging

import (
	"fmt"
)

// type PlayerStrategies interface {
// 	basic(p *Player) *Player
// }
type PlayerStrategy func(*Player) ([]string, []uint8)

var (
	countStrategies  = []PlayerStrategy{aceFiveCount, hiLoCount}
	countDescription = []string{"Ace/Five Count", "HiLo Count"}
	randStrategies   = []PlayerStrategy{basic, wizardOfOdds}
	description      = []string{"Basic", "Wizard of Odds"}
)

func randomNonCountPlayerStrategy() (PlayerStrategy, string) {
	randomInt := randInt(0, len(randStrategies))
	return randStrategies[randomInt], description[randomInt]
}
func randomCountPlayerStrategies() (PlayerStrategy, string) {
	randomInt := randInt(0, len(countStrategies))
	return countStrategies[randomInt], countDescription[randomInt]
}

//a very basic/stupid strategy, hit if below 17, stand otherwise
func basic(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}
	tableCounter := p.table.calculateTableCount()
	runningCount := p.count(tableCounter)
	trueCount := runningCount / DEFAULTDECKPERSHOE
	fmt.Printf("[counting strategy: %s, Running Count: %f, True Count: %f]\n", p.countDescription, runningCount, trueCount)
	fmt.Print("[playing strategy: basic]: ")
	if p.currentBet > 0 {
		for i := 0; i < len(p.hands); i++ {
			if value, _ := p.calculateHandValue(uint8(i)); value < 17 {
				actions = append(actions, "hit")
				handIndices = append(handIndices, uint8(i))
			} else {
				actions = append(actions, "stand")
				handIndices = append(handIndices, uint8(i))
			}
		}
	}
	return actions, handIndices
}

/*
FROM: http://wizardofodds.com/games/blackjack/strategy/4-decks/

Surrender

Surrender hard 16 (but not a pair of 8s) vs. dealer 9, 10, or A, and hard 15 vs. dealer 10.
Split

Always split aces and 8s.
Never split 5s and 10s.
Split 2s and 3s against a dealer 4-7, and against a 2 or 3 if DAS is allowed.
Split 4s only if DAS is allowed and the dealer shows a 5 or 6.
Split 6s against a dealer 3-6, and against a 2 if DAS is allowed.
Split 7s against a dealer 2-7.
Split 9s against a dealer 2-6 or 8-9.

Double hard 9 vs. dealer 3-6.
Double hard 10 except against a dealer 10 or A.
Double hard 11 except against a dealer A.
Double soft 13 or 14 vs. dealer 5-6.
Double soft 15 or 16 vs. dealer 4-6.
Double soft 17 or 18 vs. dealer 3-6.

Always hit hard 11 or less.
Stand on hard 12 against a dealer 4-6, otherwise hit.
Stand on hard 13-16 against a dealer 2-6, otherwise hit.
Always stand on hard 17 or more.
Always hit soft 17 or less.
Stand on soft 18 except hit against a dealer 9, 10, or A.
Always stand on soft 19 or more.

*/
func wizardOfOdds(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}
	tableCounter := p.table.calculateTableCount()
	runningCount := p.count(tableCounter)
	trueCount := runningCount / DEFAULTDECKPERSHOE
	fmt.Printf("[counting strategy: %s]: Running Count: %f, True Count: %f\n", p.countDescription, runningCount, trueCount)
	fmt.Print("[playing strategy: WizardOfOdds]: ")

	dealerCard := p.table.dealer.curHand.cards[0]
	for i, hand := range p.hands {
		playerHandValue, isSoft := p.calculateHandValue(uint8(i))
		curAction := wizardOfOddsActionLogic(p, dealerCard, playerHandValue, isSoft, hand)
		actions = append(actions, curAction)
		handIndices = append(handIndices, uint8(i))
	}

	return actions, handIndices

}

// prepare yourself for the longest fucking if else piece of shit ever
func wizardOfOddsActionLogic(p *Player, dealerCard *Card, playerHandValue uint8, isSoft bool, hand *Hand) string {
	if len(hand.cards) == 2 {
		if hand.cards[0].value == "8" && hand.cards[1].value == "8" && dealerCard.value == "A" {
			return "surrender"
		}
		if (hand.cards[0].value == "A" && hand.cards[1].value == "A") || (hand.cards[0].value == "8" && hand.cards[1].value == "8") {
			return "split"
		}
		if (hand.cards[0].value == "2" && hand.cards[1].value == "2") || (hand.cards[0].value == "3" && hand.cards[1].value == "3") && (dealerCard.numberValue <= 7 && dealerCard.numberValue >= 4) {
			return "split"
		}
		if (hand.cards[0].value == "6" && hand.cards[1].value == "6") && (dealerCard.numberValue <= 6 && dealerCard.numberValue >= 3) {
			return "split"
		}
		if (hand.cards[0].value == "7" && hand.cards[1].value == "7") && (dealerCard.numberValue <= 7 && dealerCard.numberValue >= 2) {
			return "split"
		}
		if (hand.cards[0].value == "9" && hand.cards[1].value == "9") && ((dealerCard.numberValue <= 6 && dealerCard.numberValue >= 2) || (dealerCard.numberValue <= 9 && dealerCard.numberValue >= 8)) {
			return "split"
		}
	}
	if (playerHandValue == 15 || playerHandValue == 17) && dealerCard.value == "A" {
		return "surrender"
	} else if (playerHandValue == 15 && !isSoft) && dealerCard.value == "10" {
		return "surrender"
	} else if (playerHandValue == 16 && !isSoft) && (dealerCard.value == "9" || dealerCard.value == "10" || dealerCard.value == "A") {
		return "surrender"
	} else if (playerHandValue == 9 && !isSoft) && (dealerCard.value == "3" || dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
		return "double"
	} else if (playerHandValue == 10 && !isSoft) && (dealerCard.value != "10" && dealerCard.value != "A") {
		return "double"
	} else if (playerHandValue == 11 && !isSoft) && dealerCard.value != "A" {
		return "double"
	} else if (playerHandValue == 13 || playerHandValue == 14 || isSoft) && (dealerCard.value == "5" || dealerCard.value == "6") {
		return "double"
	} else if (playerHandValue == 15 || playerHandValue == 16 || isSoft) && (dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
		return "double"
	} else if (playerHandValue == 17 || playerHandValue == 18 || isSoft) && (dealerCard.value == "3" || dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
		return "double"
	} else if playerHandValue <= 11 && !isSoft {
		return "hit"
	} else if playerHandValue == 12 && !isSoft && !(dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
		return "hit"
	} else if (playerHandValue == 13 || playerHandValue == 14 || playerHandValue == 15 || playerHandValue == 16) && !isSoft && !(dealerCard.value == "2" || dealerCard.value == "3" || dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
		return "hit"
	} else if playerHandValue < 17 && isSoft {
		return "hit"
	} else {
		return "stand"
	}
}

// corresponding to Hi Lo counting strategy
func hiLoCount(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}
	tableCounter := p.table.calculateTableCount()
	runningCount := p.count(tableCounter)
	trueCount := runningCount / DEFAULTDECKPERSHOE
	fmt.Printf("[counting strategy: %s, Running Count: %f, True Count: %f]\n", p.countDescription, runningCount, trueCount)
	fmt.Print("[playing strategy: HiLo Count]: ")
	dealerCard := p.table.dealer.curHand.cards[0]

	for i, hand := range p.hands {
		var curAction string
		playerHandValue, isSoft := p.calculateHandValue(uint8(i))

		if trueCount > 0  && playerHandValue < 10{
				curAction = "double"
		} else {
			curAction = wizardOfOddsActionLogic(p, dealerCard, playerHandValue, isSoft, hand)
		}
		actions = append(actions, curAction)
		handIndices = append(handIndices, uint8(i))
	}
	return actions, handIndices
}

// corresponding to Ace/Five counting strategy
func aceFiveCount(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}
	tableCounter := p.table.calculateTableCount()
	runningCount := p.count(tableCounter)
	trueCount := runningCount / DEFAULTDECKPERSHOE
	fmt.Printf("[counting strategy: %s, Running Count: %f, True Count: %f]\n", p.countDescription, runningCount, trueCount)
	fmt.Print("[playing strategy: Ace/Five Count]: ")
	dealerCard := p.table.dealer.curHand.cards[0]
	for i, hand := range p.hands {
		var curAction string
		playerHandValue, isSoft := p.calculateHandValue(uint8(i))
		if trueCount > 2 {
			curAction = "double"
		} else {
			curAction = wizardOfOddsActionLogic(p, dealerCard, playerHandValue, isSoft, hand)
		}
		actions = append(actions, curAction)
		handIndices = append(handIndices, uint8(i))

	}
	return actions, handIndices

}
