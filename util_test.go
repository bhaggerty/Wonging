package wonging

import (
	"testing"
)

func Test_CheckCardsValueEqual(t *testing.T) {
	card1 := new(Card).Initialize("2", 2, "Diamonds")
	card2 := new(Card).Initialize("3", 3, "Diamonds")
	if !checkCardsValueEqual(card1, card2) {
		t.Log("CheckCardsValueEqual() test passed")
	} else {
		t.Error("CheckCardsValueEqual() did not work as expected.")
	}
}
