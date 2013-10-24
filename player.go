package Wonging

type Player struct {
	id uint8
	//current table the player is sitting at
	tableId int8
	//current hand
	hand *Hand
	//how much is the player betting
	currentBet float32
	//how much money does the player have
	totalCash float32
}

func (p *Player) bet(money float32) {
	if money <= 0 {
		fmt.Println("Invalid bet")
	} else {
		p.currentBet += money
		p.totalCash -= money
	}
}
