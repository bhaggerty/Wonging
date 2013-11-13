package wonging

import (
	"fmt"
)

type PlayerStrategies interface {
	basic(p *Player) *Player
}

//a basic strategy
func basic(p *Player) {
	if value, _ := p.calculateHandValue(handIndex); value < 17 {
		p.hit()
	} else {
		p.stand()
	}
}
