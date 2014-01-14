package wonging

import (
	"testing"
)

func predefinePlayer() *Player {
	casino := predefineCasino()
	table := predefineTable()
	return new(Player).Initialize(1, casino, table).win(20)
}
func predefineAnotherPlayer() *Player {
	casino := predefineCasino()
	table := predefineTable()
	return new(Player).Initialize(2, casino, table).win(1)
}

func Test_InitializePlayer(t *testing.T) {
	player := predefinePlayer()
	if player.id != 1 {
		t.Error("Initialize() [id attribute] did not work as expected.")
	} else if player.casino == nil {
		t.Error("Initialize() [casino attribute] did not work as expected.")
	} else if player.table == nil {
		t.Error("Initialize() [tables attribute] did not work as expected.")
	} else if player.currentBet != 0 {
		t.Error("Initialize() [currentBet attribute] did not work as expected.")
	} else if player.totalCash != DEFAULTPLAYERSTARTINGCASH+20 {
		t.Error("Initialize() [totalCash attribute] did not work as expected.")
	} else if player.action == nil {
		t.Error("Initialize() [action attribute] did not work as expected.")
	} else if player.count == nil {
		t.Error("Initialize() [count attribute] did not work as expected.")
	} else if player.winCount != 1 {
		t.Error("Initialize() [winCount attribute] did not work as expected.")
	} else if player.loseCount != 0 {
		t.Error("Initialize() [loseCount attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}

func Test_winLose(t *testing.T) {
	player := predefinePlayer()
	player.win(30)
	if player.currentBet != 0 {
		t.Error("win() [currentBet attribute] did not work as expected.")
	} else if player.winCount != 2 {
		t.Error("win() [winCount attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}

	card := predefineCard()
	player.acceptCard(card, 0)
	player.bet(30, 0)
	player.lose()
	if player.currentBet != 0 {
		t.Error("lose() [currentBet attribute] did not work as expected.")
	} else if player.loseCount != 1 {
		t.Error("lose() [loseCount attribute] did not work as expected.")
	} else {
		t.Log("Initialize() test passed")
	}
}
