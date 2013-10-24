package Wonging

type Player struct {
	id         uint8
	tableId    int8
	hand       *Hand
	currentBet float32
	totalCash  float32
}
