//Game object, to track results for game and prepare for
//analytics later
package wonging

import (
	"fmt"
)

type Game struct {
	//pointer to parent
	table *Table
	//casino earning
	casinoEarning float64
	//an array of money in/out for players
	playerResult []float64
}

func (g *Game) Initialize(t *Table) *Game {
	g.table = t
	g.casinoEarning = 0
	//player betting amount
	for i := 0; i < len(t.players); i++ {
		g.playerResult = append(g.playerResult, 0)
	}
	return g
}

func (g *Game) updatePlayerResult(t *Table) *Game {
	var sum float64 = g.casinoEarning
	for i := 0; i < len(t.players); i++ {
		g.playerResult[i] = t.players[i].profit()
		sum -= g.playerResult[i]
	}
	g.casinoEarning = sum
	return g
}

func (g *Game) biggestWinner() int {
	return MaxFloatS(g.playerResult)
}
func (g *Game) biggestLoser() int {
	return MinFloatS(g.playerResult)
}

func (g *Game) PrintGame() {
	fmt.Print(g.Description())
}

func (g *Game) Description() string {
	return fmt.Sprintf(">>  Game Result:\n    players: %@\n    casino: %.2f\n", g.playerResult, g.casinoEarning)
}

func (g *Game) HTMLString() string {
	html := "<div class=\"game\">Game results:  Players: ["
	for _, playerResult := range g.playerResult {
		html += fmt.Sprintf("%.2f, ", playerResult)
	}
	html += fmt.Sprintf("]   Casino:%.2f</div>", g.casinoEarning)
	return html
}
