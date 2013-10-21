package wonging

import (
	"github.com/josephyzhou/wonging"
	"strconv"
	"testing"
)

//generate a "busted" hand
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

//generate a blackjack hand
func predefineBJ() *wonging.Hand {
	card1 := new(wonging.Card).NewCard("5", int8(5), "Clubs")
	card2 := new(wonging.Card).NewCard("6", int8(6), "Clubs")
	card3 := new(wonging.Card).NewCard("10", int8(10), "Clubs")
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

func Test_DetermineOutcome(t *testing.T) {
	handBJ := predefineBJ()
	var outcome string
	outcome = handBJ.DetermineOutcome(22)
	if outcome != "Dealer busted" {
		t.Error("DetermineOutcome() - Expected 'Dealer busted', Actual: " + outcome)
	} else {
		t.Log("DetermineOutcome() ifBusted case works")
	}
	outcome = handBJ.DetermineOutcome(10)
	if outcome != "Player wins" {
		t.Error("DetermineOutcome() - Expected 'Player wins', Actual: " + outcome)
	} else {
		t.Log("DetermineOutcome() PlayerWins case works")
	}

	//trying a case passing in two parameters
	outcome = handBJ.DetermineOutcome(15, 22)
	if outcome != "Player busted" {
		t.Error("DetermineOutcome() - Expected 'Player busted', Actual: " + outcome)
	} else {
		t.Log("DetermineOutcome() PlayerBusted case works")
	}
}
