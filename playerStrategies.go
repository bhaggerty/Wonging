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
type PlayerStrategy func(*Player) (string, uint8)

func randomPlayerStrategy() PlayerStrategy {
	strategies := []PlayerStrategy{basic}
	return strategies[randInt(0, len(strategies))]
}

//a very basic/stupid strategy, hit if below 17, stand otherwise
func basic(p *Player) (string, uint8) {
	fmt.Print("[strategy: basic]: ")
	if p.currentBet > 0 {
		for i := 0; i < len(p.hands); i++ {
			if value, _ := p.calculateHandValue(uint8(i)); value < 17 {
				return "hit", uint8(i)
			} else {
				return "stand", uint8(i)
			}
		}
	}
	return "stand", 0
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
Double

Double hard 9 vs. dealer 3-6.
Double hard 10 except against a dealer 10 or A.
Double hard 11 except against a dealer A.
Double soft 13 or 14 vs. dealer 5-6.
Double soft 15 or 16 vs. dealer 4-6.
Double soft 17 or 18 vs. dealer 3-6.
Hit or Stand

Always hit hard 11 or less.
Stand on hard 12 against a dealer 4-6, otherwise hit.
Stand on hard 13-16 against a dealer 2-6, otherwise hit.
Always stand on hard 17 or more.
Always hit soft 17 or less.
Stand on soft 18 except hit against a dealer 9, 10, or A.
Always stand on soft 19 or more.
As I've said many times, the above strategy will be fine under any set of rules. However, for you perfectionists out there, here are the modifications to make if the dealer hits a soft 17.

Surrender 15, a pair of 8s, and 17 vs. dealer A.
Double 11 vs. dealer A.
Double soft 18 vs. dealer 2.
Double soft 19 vs. dealer 6.
*/
func wizardOfOzz(p *Player) (string, uint8) {

}
