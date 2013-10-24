package wonging

// -1 means cannot find
func checkDealerContain(d *Dealer, dealers []*Dealer) int {
	index := -1
	for i := 0; i < len(dealers); i++ {
		if d.id == dealers[i].id {
			index = i
		}
	}
	return index
}

func checkPlayerContain(p *Player, players []*Player) int {
	index := -1
	for i := 0; i < len(players); i++ {
		if p.id == players[i].id {
			index = i
		}
	}
	return index
}
