package wonging

import (
	"testing"
)

func Test_InitializeGame(t *testing.T) {
	table := new(Table).Initialize(1, new(Casino).Initialize(0))
	game := new(Game).Initialize(table)
	if game.table == nil {
		t.Error("Initialize() [table attribute] did not work as expected.")
	} else if game.casinoEarning != 0 {
		t.Error("Initialize() [casinoEarning attribute] did not work as expected.")
	} else if !(len(table.players) == 0 && game.playerResult == nil) {
		t.Error("Initialize() [playerResult attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}

func Test_updatePlayerResult(t *testing.T) {

}
