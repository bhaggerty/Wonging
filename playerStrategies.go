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

func randomPlayerStrategy() PlayerStrategy {
	strategies := []PlayerStrategy{basic, wizardOfOzz}
	return strategies[randInt(0, len(strategies))]
}

//a very basic/stupid strategy, hit if below 17, stand otherwise
func basic(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}

	fmt.Print("[strategy: basic]: ")
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
As I've said many times, the above strategy will be fine under any set of rules. However, for you perfectionists out there, here are the modifications to make if the dealer hits a soft 17.

*/
func wizardOfOzz(p *Player) ([]string, []uint8) {
	actions := []string{}
	handIndices := []uint8{}

	fmt.Print("[strategy: WizardOfOzz]: ")

	// prepare yourself for the longest fucking if else piece of shit ever
	dealerCard := p.table.dealer.curHand.cards[0]
	for i, hand := range p.hands {
		playerHandValue, isSoft := p.calculateHandValue(uint8(i))
		var curAction string
		if (playerHandValue == 15 || playerHandValue == 17) && dealerCard.value == "A" {
			curAction = "surrender"
		} else if (playerHandValue == 15 && !isSoft) && dealerCard.value == "10" {
			curAction = "surrender"
		} else if (playerHandValue == 16 && !isSoft) && (dealerCard.value == "9" || dealerCard.value == "10" || dealerCard.value == "A") {
			curAction = "surrender"
		} else if (playerHandValue == 9 && !isSoft) && (dealerCard.value == "3" || dealerCard.value == "4" || dealerCard.value == "5" || dealerCard.value == "6") {
			curAction = "double"
		} else if (playerHandValue == 10 && !isSoft) && (dealerCard.value != "10" && dealerCard.value != "A") {
			curAction = "double"
		} else if (playerHandValue == 11 && !isSoft) && dealerCard.value != "A" {
			curAction = "double"
		} else if (playerHandValue == 13 && playerHandValue == 14 && isSoft) && (dealerCard.value == "5" && dealerCard.value == "6") {
			curAction = "double"
		} else if (playerHandValue == 15 && playerHandValue == 16 && isSoft) && (dealerCard.value == "4" && dealerCard.value == "5" && dealerCard.value == "6") {
			curAction = "double"
		} else if (playerHandValue == 17 && playerHandValue == 18 && isSoft) && (dealerCard.value == "3" && dealerCard.value == "4" && dealerCard.value == "5" && dealerCard.value == "6") {
			curAction = "double"
		} else if playerHandValue <= 11 && !isSoft {
			curAction = "hit"
		} else if playerHandValue < 17 && isSoft {
			curAction = "hit"
		} else {
			curAction = "stand"
		}

		if len(hand.cards) == 2 {
			if hand.cards[0].value == "8" && hand.cards[1].value == "8" && dealerCard.value == "A" {
				curAction = "surrender"
			} else if (hand.cards[0].value == "A" && hand.cards[1].value == "A") || (hand.cards[0].value == "8" && hand.cards[1].value == "8") {
				curAction = "split"
			} else if (hand.cards[0].value == "2" && hand.cards[1].value == "2") || (hand.cards[0].value == "3" && hand.cards[1].value == "3") && (dealerCard.numberValue <= 7 && dealerCard.numberValue >= 4) {
				curAction = "split"
			} else if (hand.cards[0].value == "6" && hand.cards[1].value == "6") && (dealerCard.numberValue <= 6 && dealerCard.numberValue >= 3) {
				curAction = "split"
			} else if (hand.cards[0].value == "7" && hand.cards[1].value == "7") && (dealerCard.numberValue <= 7 && dealerCard.numberValue >= 2) {
				curAction = "split"
			} else if (hand.cards[0].value == "9" && hand.cards[1].value == "9") && ((dealerCard.numberValue <= 6 && dealerCard.numberValue >= 2) || (dealerCard.numberValue <= 9 && dealerCard.numberValue >= 8)) {
				curAction = "split"
			}
		}

		actions = append(actions, curAction)
		handIndices = append(handIndices, uint8(i))
	}

	return actions, handIndices

}
