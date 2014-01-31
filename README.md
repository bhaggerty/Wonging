Wonging
=======
[![GoDoc](https://godoc.org/github.com/josephyzhou/Wonging?status.png)](https://godoc.org/github.com/josephyzhou/Wonging)
[![Build Status](https://travis-ci.org/josephyzhou/Wonging.png?branch=master)](https://https://travis-ci.org/josephyzhou/Wonging)
####(WORK IN PROGRESS!)
## Table of Contents ##
- [What](#what-is-it)
- [Plan](#plan)
- [Install](#install)
- [Testing](#testing)
- [Explaination](#explaination)
- [License](#license)


### What is it ###
A little card counting tool/simulator (for educational purpose only!)

Wonging: Back-counting, consists of standing behind a blackjack table that other players are playing on, and counting the cards as they are dealt. Stanford Wong first proposed the idea of back-counting, and the term "Wong" comes from his name.

### Plan ###
The plan is to create a complete simulator of the backjack table section of a typical casino, from the random # of players per table, to randomized counting strategies each player use or not use. The goal is to have a bird-eye view on casino operations and see how much each player would lose/win base on different strategies used. The simulator will also calculate winning chance for each hand for each player, in order to statistically determine when is the right time to start joining a game or leave a game as a player

The second phase of the project will cover group counting/strategies, by combining players into groups, this simulator hopefully may reveal different(better) winning percentages than individual operations

Third phase? Sims Casino - AI for the game

### Install ###
installing GoLang:
http://golang.org/doc/install

`go get github.com/josephyzhou/Wonging`


###Testing

`go test -v`

### Explaination ###
`cd /simulator`

`go run simulator.go`

####The Console Output ####
Console output is separated into 3 sections, corresponding to this piece of code in the simulator.go
```
casino := new(wonging.Casino).Initialize(0)
	casino.DistributeDealers()
	casino.DistributePlayers()
	casino.PrintCasino()

	//casino fully operational
	casino.Start()
	casino.PrintCasino()
	casino.Log()
	casino.GenerateHTMLMap()
```
1. First `PrintCasino()` will print the casino in its initial state - before all the game simulations. You should see the dealers/players being distributed, how many dealers/players are idle after the initial distribution, and the amount of money different entities have.
2. Immediately following, you should see the game simulations. The default setting (in `config.go`) is to run 200 games. So `casino.Start()` would start the game, and you should see the console output
3. At the end, we call `PrintCasino()` again to print the end state of the casino. This is interesting because we will see the total profit of casino, each player's performance.

#### The HTML Output ####
I've just started this. Right now `casino.GenerateHTMLMap()` should generate a new folder `/html` and in there you should be able to see MANY different .html files. Click on `casino0.html` and you should be able to view the end state of the casino this way. The plan is to have one html with cool js/css/graphics

###License
MIT - see LICENSE file
