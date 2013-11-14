package wonging

import (
// "fmt"
)

type PlayerStrategies interface {
	basic(p *Player) *Player
}

//a basic strategy
func basic(p *Player) {
	for i := 0; i < len(p.hands); i++ {
		if value, _ := p.calculateHandValue(uint8(i)); value < 17 {
			p.hit(uint8(i))
		} else {
			p.stand(uint8(i))
		}
	}
}

func doubleDown() {
}
