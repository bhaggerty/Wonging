package main

import (
	"github.com/josephyzhou/wonging"
)

func main() {

	//setting up
	casino := new(wonging.Casino).Initialize(0)
	casino.DistributeDealers()
	casino.DistributePlayers()
	casino.PrintCasino()

	//casino fully operational
	casino.Start()
	casino.PrintCasino()
}
