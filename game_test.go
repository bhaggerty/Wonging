package wonging

import (
	"testing"
)

func predefineGame() *Game {
	table := predefineTable()
	return new(Game).Initialize(table)
}

func predefineGameWithPlayerResults() *Game {
	table := predefineTableWithPlayers()
	if table != nil {
		game := new(Game).Initialize(table)
		game.updatePlayerResult(table)
		return game
	}
	return nil
}

func Test_InitializeGame(t *testing.T) {
	game := predefineGame()
	if game.table == nil {
		t.Error("Initialize() [table attribute] did not work as expected.")
	} else if game.casinoEarning != 0 {
		t.Error("Initialize() [casinoEarning attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}

func Test_UpdatePlayerResult(t *testing.T) {
	game := predefineGameWithPlayerResults()
	if game.playerResult[0] == 20 && game.playerResult[1] == 1 {
		t.Log("UpdatePlayerResult() test passed")
	} else {
		t.Error("UpdatePlayerResult() did not work as expected.")
	}
}

func Test_biggestWinnerLoser(t *testing.T) {
	game := predefineGameWithPlayerResults()
	if game.biggestWinner() == 0 {
		t.Log("biggestWinner() test passed")
	} else {
		t.Error("biggestWinner() did not work as expected.")
	}
	if game.biggestLoser() == 1 {
		t.Log("biggestLoser() test passed")
	} else {
		t.Error("biggestLoser() did not work as expected. ", game.biggestLoser())
	}
}
