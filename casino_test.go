package wonging

import (
	"testing"
)

func predefineCasino() *Casino {
	return new(Casino).Initialize(1)
}

func Test_InitializeCasino(t *testing.T) {
	casino := predefineCasino()
	if casino.id != 1 {
		t.Error("Initialize() [id attribute] did not work as expected.")
	} else if casino.bank != DEFAULTCASINOSTARTINGCASH {
		t.Error("Initialize() [bank attribute] did not work as expected.")
	} else if casino.tables == nil || len(casino.tables) != DEFAULTNUMBEROFTABLESPERCASINO {
		t.Error("Initialize() [tables attribute] did not work as expected.")
	} else if casino.idleDealers == nil || len(casino.idleDealers) != DEFAULTNUMBEROFDEALERSPERCASINO {
		t.Error("Initialize() [idleDealers attribute] did not work as expected.")
	} else if casino.idlePlayers == nil || len(casino.idlePlayers) != DEFAULTNUMBEROFPLAYERSPERCASINO {
		t.Error("Initialize() [idlePlayers attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}

func Test_winLoseProfit(t *testing.T) {
	casino := predefineCasino()
	casino.win(15)

	if casino.bank != DEFAULTCASINOSTARTINGCASH+15 {
		t.Error("win() did not work as expected.")
	} else {
		t.Log("win() test passed")
	}

	casino.lose(2)

	if casino.bank != DEFAULTCASINOSTARTINGCASH+13 {
		t.Error("lose() did not work as expected.")
	} else {
		t.Log("lose() test passed")
	}

	if casino.totalProfit() != 13 {
		t.Error("totalProfit() did not work as expected.")
	} else {
		t.Log("totalProfit() test passed")
	}
}

func Test_distributeDealersPlayers(t *testing.T) {
	casino := predefineCasino()
	casino.DistributeDealers()
	if DEFAULTNUMBEROFDEALERSPERCASINO != len(casino.tables) {
		if casino.totalActiveDealers() != uint8(len(casino.tables)) && casino.totalInactiveDealers() != uint8(DEFAULTNUMBEROFDEALERSPERCASINO-len(casino.tables)) {
			t.Error("DistributedDealers() [diff number of tables/dealers] did not work as expected.")
		} else {
			t.Log("DistributedDealers() [diff number of tables/dealers] test passed")
		}
	} else {
		if casino.totalActiveDealers() != uint8(DEFAULTNUMBEROFDEALERSPERCASINO) && casino.totalInactiveDealers() != 0 {
			t.Error("DistributedDealers() [equal number of tables/dealers] did not work as expected.")
		} else {
			t.Log("DistributedDealers() [equal number of tables/dealers] test passed")
		}
	}

	// casino.DistributePlayers()
	// if DEFAULTNUMBEROFPLAYERSPERCASINO != len(casino.tables) {
	// 	if casino.totalActivePlayers() != len(casino.tables) && casino.totalInactivePlayers() != DEFAULTNUMBEROFPLAYERSPERCASINO-len(casino.tables) {
	// 		t.Error("DistributedDealers() [diff number of tables/dealers] did not work as expected.")
	// 	} else {
	// 		t.Log("DistributedDealers() [diff number of tables/dealers] test passed")
	// 	}
	// }
}
