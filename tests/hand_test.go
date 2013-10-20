package wonging

import (
	"github.com/josephyzhou/wonging"
	"testing"
)

func Test_calculateValue(t *testing.T) {
	card1 := new(wonging.Card).NewCard("9", int8(9), "Clubs")
	card2 := new(wonging.Card).NewCard("J", int8(10), "Hearts")
	card3 := new(wonging.Card).NewCard("K", int8(10), "Diamonds")
	hand := new(wonging.Hand)
	hand.AddCard(card1)
	hand.AddCard(card2)
	hand.AddCard(card3)
	totalValue := hand.CalculateValue()
	if totalValue != 29 {
		t.Error("CalculateValue() did not work as expected.")
	} else {
		t.Log("CalculateValue() test passed")
	}
}
