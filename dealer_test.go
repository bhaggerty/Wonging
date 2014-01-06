package wonging

import (
	"testing"
)

func preDefindDealer() *Dealer {
	table := new(Table).Initialize(1, new(Casino).Initialize(0))
	shoe := new(Deck).Initialize(2)
	return new(Dealer).Initialize(1, table, shoe)
}

func Test_InitializeDealer(t *testing.T) {
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

func Test_calculateHandValue(t *testing.T) {
	dealer := preDefindDealer()
	dealer.curHand = predefineHand()
	dealer.faceDown = new(Card).Initialize("A", 1, "Diamonds")

	//test for calculate entire hand
	handValue, isSoft := dealer.calculateHandValue()
	if handValue != 30 || isSoft {
		t.Error("calculateHandValue() did not work as expected.")
	} else {
		t.Log("calculateHandValue() test passed")
	}

	//test for calculate visible hands only
	visHandValue, isSoft := dealer.calculateVisibleHandValue()
	if visHandValue != 29 || isSoft {
		t.Error("calculateVisibleHandValue() did not work as expected.")
	} else {
		t.Log("calculateVisibleHandValue() test passed")
	}
}

func Test_isBusted(t *testing.T) {
	dealer := preDefindDealer()
	dealer.curHand = predefineHand()
	dealer.faceDown = new(Card).Initialize("A", 1, "Diamonds")
	if !(dealer.isBusted()) {
		t.Error("isBusted() did not work as expected.")
	} else {
		t.Log("isBusted() test passed")
	}
}

func Test_dealerActions(t *testing.T) {
	dealer := preDefindDealer()

	//test for deal
	card := dealer.deal()
	if card == nil {
		t.Error("deal() did not work as expected.")
	} else {
		t.Log("deal() test passed")
	}

	//test for deal self
	dealer.dealSelf()
	if dealer.faceDown == nil || len(dealer.curHand.cards) > 0 {
		t.Error("dealSelf() [1 card case] did not work as expected.")
	} else {
		t.Log("dealSelf() [1 card case] test passed")
	}

	dealer.dealSelf()
	if dealer.faceDown == nil || len(dealer.curHand.cards) != 1 {
		t.Error("dealSelf() [2 card case] did not work as expected.")
	} else {
		t.Log("dealSelf() [2 card case] test passed")
	}

	// test for deck reset
	dealer.resetDeck()
	if len(dealer.shoe.cards) != 52*DEFAULTDECKPERSHOE {
		t.Error("resetDeck() did not work as expected.")
	} else {
		t.Log("resetDeck() test passed")
	}
}
