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
