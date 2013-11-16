package wonging

// type PlayerStrategies interface {
// 	basic(p *Player) *Player
// }
type PlayerStrategy func(*Player) string

func randomPlayerStrategy() PlayerStrategy {
	strategies := []PlayerStrategy{basic}
	return strategies[randInt(0, len(strategies)-1)]
}

//a basic strategy
func basic(p *Player) string {
	for i := 0; i < len(p.hands); i++ {
		if value, _ := p.calculateHandValue(uint8(i)); value < 17 {
			return "hit"
		} else {
			return "stand"
		}
	}
	return ""
}

func doubleDown() {
}
