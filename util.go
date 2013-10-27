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

func combineCounters(counters []*Counter) *Counter {
	combinedCounter := new(Counter).initialize()
	for _, counter := range counters {
		combinedCounter.HiLo += counter.HiLo
		combinedCounter.HiOpt1 += counter.HiOpt1
		combinedCounter.HiOpt2 += counter.HiOpt2
		combinedCounter.KO += counter.KO
		combinedCounter.Omega2 += counter.Omega2
		combinedCounter.Red7 += counter.Red7
		combinedCounter.ZenCount += counter.ZenCount
	}
	return combinedCounter
}

func checkCardsValueEqual(c1 *Card, c2 *Card) bool {
	return c1.value == c2.value
}
