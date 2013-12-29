package wonging

import (
	"testing"
)

func Test_Init(t *testing.T) {
	shoe := new(Deck).Initialize(2)
	if len(shoe.cards) != 104 {
		t.Error("Initialize() did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}

func Test_PopRandom(t *testing.T) {
	shoe := new(Deck).Initialize(1)
	shoe.popRandom()
	if len(shoe.cards) != 51 {
		t.Error("PopRandom() did not work as expected.")
	} else {
		t.Log("PopRandom() test passed")
	}
}

func Test_Pop(t *testing.T) {
	shoe := new(Deck).Initialize(1)
	poppedCard := shoe.pop()
	if len(shoe.cards) != 51 || checkShoeContain(poppedCard, shoe) != -1 {
		t.Error("Pop() did not work as expected.")
	} else {
		t.Log("Pop() test passed")
	}
}
