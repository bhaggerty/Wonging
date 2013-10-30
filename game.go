//Game object, to track results for game and prepare for
//analytics later
package wonging

import (
	"fmt"
)

type Game struct {
	//an array of money in/out for both dealers and players
	//index 0 is dealer
	//index 1... are players, tracking players array position
	moneyResult []float32

	//current round or final when everythings done
	round uint8
}

func (g *Game) Initialize() *Game {
	g.round = 1
	return g
}

func (g *Game) biggestWinner() int {
	return MaxFloatS(g.moneyResult)
}
func (g *Game) biggestLoser() int {
	return MinFloatS(g.moneyResult)
}

func (g *Game) PrintGame() {
	fmt.Printf(">> Game Result: %@", g.moneyResult)
}
