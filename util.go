package wonging

import (
	"math/rand"
	"time"
)

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

func MinFloatS(v []float32) int {
	var index int
	var m float32
	if len(v) > 0 {
		m = v[0]

	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
			index = i
		}
	}
	return index
}

func MaxFloatS(v []float32) int {
	var index int
	var m float32
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
			index = i
		}
	}
	return index
}

func randInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min == max {
		return min
	} else {
		return min + rand.Intn(max-min)
	}
}
