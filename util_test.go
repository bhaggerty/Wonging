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

func Test_CheckPlayerContain(t *testing.T) {
	casino := new(Casino).Initialize(1)
	table := new(Table).Initialize(1, casino)
	p1 := new(Player).Initialize(1, casino, table)
	p2 := new(Player).Initialize(2, casino, table)
	p3 := new(Player).Initialize(3, casino, table)
	players := append([]*Player{}, p1, p2, p3)
	if checkPlayerContain(p1, players) == -1 {
		t.Error("checkPlayerContain() did not work as expected.")
	} else {
		t.Log("checkPlayerContain() test passed")
	}
}

func Test_CheckDealerContain(t *testing.T) {
	casino := new(Casino).Initialize(1)
	shoe := new(Deck).Initialize(1)
	table := new(Table).Initialize(1, casino)
	d1 := new(Dealer).Initialize(1, table, shoe)
	d2 := new(Dealer).Initialize(2, table, shoe)
	d3 := new(Dealer).Initialize(3, table, shoe)
	dealers := append([]*Dealer{}, d1, d2, d3)
	if checkDealerContain(d1, dealers) == -1 {
		t.Error("CheckDealerContain() did not work as expected.")
	} else {
		t.Log("CheckDealerContain() test passed")
	}
}
