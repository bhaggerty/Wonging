//Game object, to track results for game and prepare for
//analytics later
package wonging

import (
	"fmt"
)

type Game struct {
	//pointer to parent
	table *Table

	//an array of money in/out for both dealers and players
	//index 0 is dealer
	//index 1... are players, tracking players array position
	moneyResult []float64

	//current round or final when everythings done
	round uint8
}

func (g *Game) Initialize(t *Table) *Game {
	g.table = t
	g.round = 1
	//dealer
	g.moneyResult = append(g.moneyResult, 0)
	//player betting amount
	for i := 1; i < len(t.players); i++ {
		g.moneyResult = append(g.moneyResult, 0-t.players[i-1].currentBet)
	}
	return g
}

func (g *Game) updateMoneyResult() *Game {
	// g.moneyResult =
	return g
}

func (g *Game) biggestWinner() int {
	return MaxFloatS(g.moneyResult)
}
func (g *Game) biggestLoser() int {
	return MinFloatS(g.moneyResult)
}

func (g *Game) PrintGame() {
	fmt.Printf(">> Game Result:\n    round: %d\n    money: %@\n", g.round-1, g.moneyResult)
}
