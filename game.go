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

	//current round or final when everythings done
	round uint8
}

func (g *Game) Initialize(t *Table) *Game {
	g.table = t
	g.round = 0
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
	fmt.Printf(">> Game Result:\n    round: %d\n    players: %@\n    casino: %f\n", g.round, g.playerResult, g.casinoEarning)
}
