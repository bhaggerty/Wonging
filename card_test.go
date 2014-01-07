package wonging

import (
	"testing"
)

func predefineCard() *Card {
	return new(Card).Initialize("10", 10, "Diamonds")
}
func Test_Symbol(t *testing.T) {
	card := predefineCard()
	if !(card.symbol == RedText("♦")) {
		t.Error("Symbol case switch did not work as expected.")
	} else {
		t.Log("Symbol case switch test passed")
	}
}
