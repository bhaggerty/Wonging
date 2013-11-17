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

//a basic strategy, hit if below 17
func basic(p *Player) (string, uint8) {
	fmt.Print("[strategy: basic]: ")

	for i := 0; i < len(p.hands); i++ {
		if value, _ := p.calculateHandValue(uint8(i)); value < 17 {
			return "hit", uint8(i)
		} else {
			return "stand", uint8(i)
		}
	}
	return "stand", 0
}

func doubleDown() {
}
