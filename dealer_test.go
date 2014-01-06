package wonging

import (
	"testing"
)

func Test_Initialize(t *testing.T) {
	table := new(Table).Initialize(1, new(Casino).Initialize(0))
	shoe := new(Deck).Initialize(2)
	dealer := new(Dealer).Initialize(1, table, shoe)
	if dealer.id != 1 {
		t.Error("Initialize() [id attribute] did not work as expected.")
	} else if dealer.table == nil {
		t.Error("Initialize() [table attribute] did not work as expected.")
	} else if dealer.shoe == nil {
		t.Error("Initialize() [shoe attribute] did not work as expected.")
	} else if dealer.action == nil {
		t.Error("Initialize() [action attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}
