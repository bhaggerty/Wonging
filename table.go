package Wonging

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Table struct {
	//current count for all cards on table
	count int8

	//pot money in middle
	pot float32

	//total value of table
	totalCash float32

	//assuming one deck of card per dealer
	shoe *Deck

	//Many players sitting at the table
	hands []*Hand
}

func (t *Table) GetTablePlayerNumber() {
	return len(t.hands)
}
