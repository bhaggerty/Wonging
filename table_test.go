package wonging

import (
// "testing"
)

func predefineTable() *Table {
	return new(Table).Initialize(1, predefineCasino())
}

func predefineTableWithPlayer() *Table {
	table := predefineTable()
	player := predefinePlayer()
	if table.addPlayer(player) {
		return table
	}
	return nil
}

func predefineTableWithPlayers() *Table {
	table := predefineTable()
	if table.addPlayer(predefinePlayer()) && table.addPlayer(predefineAnotherPlayer()) {
		return table
	}
	return nil
}
