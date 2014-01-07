package wonging

import (
// "testing"
)

func predefinePlayer() *Player {
	casino := predefineCasino()
	table := predefineTable()
	return new(Player).Initialize(1, casino, table).win(20)
}
func predefineAnotherPlayer() *Player {
	casino := predefineCasino()
	table := predefineTable()
	return new(Player).Initialize(2, casino, table).win(1)
}
