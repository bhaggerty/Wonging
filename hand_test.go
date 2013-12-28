package wonging

import (
	"testing"
)

//generate a "busted" hand
func predefineHand() *Hand {
	card1 := new(Card).Initialize("9", uint8(9), "Clubs")
	card2 := new(Card).Initialize("J", uint8(10), "Hearts")
	card3 := new(Card).Initialize("K", uint8(10), "Diamonds")
	hand := new(Hand)
	hand.AddCard(card1).AddCard(card2).AddCard(card3)
	return hand
}

//generate a blackjack hand
func predefineBJ() *Hand {
	card1 := new(Card).Initialize("5", uint8(5), "Clubs")
	card2 := new(Card).Initialize("6", uint8(6), "Clubs")
	card3 := new(Card).Initialize("10", uint8(10), "Clubs")
	hand := new(Hand)
	hand.AddCard(card1).AddCard(card2).AddCard(card3)
	return hand
}

func Test_CalculateValue(t *testing.T) {
	hand := predefineHand()
	totalValue, isSoft := hand.CalculateValue()
	if totalValue != 29 || isSoft {
		t.Error("CalculateValue() did not work as expected.")
	} else {
		t.Log("CalculateValue() test passed")
	}
}

func Test_Description(t *testing.T) {
	hand := predefineHand()
	description := hand.Description()
	if description != "==> hand: (hard 29)\n" {
		t.Error("Description() did not work as expected.")
	} else {
		t.Log("Description() test passed")
	}
}

func Test_IsBusted(t *testing.T) {
	hand := predefineHand()
	if !hand.isBusted() {
		t.Error("isBusted() did not work as expected.")
	} else {
		t.Log("isBusted() test passed")
	}
}

// func Test_CalculateCount(t *testing.T) {
// 	hand := predefineHand()
// 	counter := hand.CalculateCount()
// 	if counter.HiLo != -2 {
// 		t.Error("CalculateCount() - Expected: -2, Actual: " + strconv.Itoa(int(counter.HiLo)))
// 	} else {
// 		t.Log("CalculateCount() HiLo test passed")
// 	}
// }

// func Test_DetermineOutcome(t *testing.T) {
// 	handBJ := predefineBJ()
// 	var outcome string
// 	outcome = handBJ.DetermineOutcome(22)
// 	if outcome != "Dealer busted" {
// 		t.Error("DetermineOutcome() - Expected 'Dealer busted', Actual: " + outcome)
// 	} else {
// 		t.Log("DetermineOutcome() isBusted case works")
// 	}
// 	outcome = handBJ.DetermineOutcome(10)
// 	if outcome != "Player wins" {
// 		t.Error("DetermineOutcome() - Expected 'Player wins', Actual: " + outcome)
// 	} else {
// 		t.Log("DetermineOutcome() PlayerWins case works")
// 	}

// 	//trying a case passing in two parameters
// 	outcome = handBJ.DetermineOutcome(15, 22)
// 	if outcome != "Player busted" {
// 		t.Error("DetermineOutcome() - Expected 'Player busted', Actual: " + outcome)
// 	} else {
// 		t.Log("DetermineOutcome() PlayerBusted case works")
// 	}
// }
