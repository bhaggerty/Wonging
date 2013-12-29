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

func Test_MinMaxFloatS(t *testing.T) {
	v := []float64{1.2, 2.4, 3.6}
	if MinFloatS(v) != 0 {
		t.Error("MinFloatS() did not work as expected.")
	} else {
		t.Log("MinFloatS() test passed")
	}

	if MaxFloatS(v) != 2 {
		t.Error("MaxFloatS() did not work as expected.")
	} else {
		t.Log("MaxFloatS() test passed")
	}

}
