package wonging

import (
	"github.com/josephyzhou/wonging"
	"strconv"
	"testing"
)

func predefineHand() *wonging.Hand {
	card1 := new(wonging.Card).NewCard("9", int8(9), "Clubs")
	card2 := new(wonging.Card).NewCard("J", int8(10), "Hearts")
	card3 := new(wonging.Card).NewCard("K", int8(10), "Diamonds")
	hand := new(wonging.Hand)
	hand.AddCard(card1)
	hand.AddCard(card2)
	hand.AddCard(card3)
	return hand
}

func Test_CalculateValue(t *testing.T) {
	hand := predefineHand()
	totalValue := hand.CalculateValue()
	if totalValue != 29 {
		t.Error("CalculateValue() did not work as expected.")
	} else {
		t.Log("CalculateValue() test passed")
	}
}

func Test_CalculateCount(t *testing.T) {
	hand := predefineHand()
	counter := hand.CalculateCount()
	if counter.HiLo != -2 {
		t.Error("CalculateCount() - Expected: -2, Actual: " + strconv.Itoa(int(counter.HiLo)))
	} else {
		t.Log("CalculateCount() HiLo test passed")
	}
}
