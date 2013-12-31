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

func Test_CombineCounters(t *testing.T) {
	counter1 := new(Counter).initialize()
	counter2 := new(Counter).initialize()
	// count some stuff
	counter1.count("A")
	counter2.count("K")

	counters := append([]*Counter{}, counter1, counter2)
	combinedCounter := combineCounters(counters)
	correctToString := "== Counter == HiLo:-2, HiOpt1:-1, HiOpt2:-2, KO:-2, Omega2:-2, Red7:-2.00, ZenCount:-3, AceFive:0"
	if combinedCounter.Description() != correctToString {
		t.Error("CombineCounters() did not work as expected.")
		t.Error(combinedCounter.Description())
	} else {
		t.Log("CombineCounters() test passed")
	}
}

func Test_CheckShoeContain(t *testing.T) {
	shoe := new(Deck).Initialize(1)
	aceOfDiamonds := new(Card).Initialize("A", 1, "Diamonds")
	if checkShoeContain(aceOfDiamonds, shoe) != -1 {
		t.Log("CheckShoeContain() test passed")
	} else {
		t.Error("CheckShoeContain() did not work as expected.")
	}
}
