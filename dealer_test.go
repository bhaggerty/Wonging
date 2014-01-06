package wonging

import (
	"testing"
)

func preDefindDealer() *Dealer {
	table := new(Table).Initialize(1, new(Casino).Initialize(0))
	shoe := new(Deck).Initialize(2)
	return new(Dealer).Initialize(1, table, shoe)
}

func Test_Initialize(t *testing.T) {
	dealer := preDefindDealer()
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

func Test_fullHand(t *testing.T) {
	dealer := preDefindDealer()
	//no card yet
	if dealer.fullHand() == nil {
		t.Log("fullHand() [no card case] test passed")
	} else {
		t.Error("fullHand() [no card case] did not work as expected.")
	}

	dealer.dealSelf()
	dealer.dealSelf()
	//should now have two cards
	if len(dealer.fullHand().cards) != 2 {
		t.Error("fullHand() [2 cards case] did not work as expected.")
	} else {
		t.Log("fullHand() [2 cards case] test passed")
	}
}
