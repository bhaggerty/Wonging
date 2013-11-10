package wonging

import (
	"fmt"
)

type PlayerStrategies interface {
	basic(p *Player) *Player
}

//a basic strategy
func basic(p *Player) {
	fmt.Println("basic")
}
