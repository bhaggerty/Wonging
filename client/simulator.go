package main

import (
	"github.com/josephyzhou/wonging"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	//setting up
	casino := new(wonging.Casino).Initialize(0)
	casino.DistributeDealers()
	casino.DistributePlayers()
	casino.PrintCasino()

	//casino fully operational
	casino.Start()
	casino.PrintCasino()
}
