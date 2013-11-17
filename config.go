package wonging

// default blackjack value
const BLACKJACK uint8 = 21

//================== Casino ==================\\
//how much money casino opens with - 1 billion
const DEFAULTCASINOSTARTINGCASH float64 = 1000000000

//how many blackjack tables are there in this casino
const DEFAULTNUMBEROFTABLESPERCASINO int = 10
const DEFAULTNUMBEROFDEALERSPERCASINO int = 12
const DEFAULTNUMBEROFPLAYERSPERCASINO int = 60

//================== Table ==================\\
// how many players can sit at one table
const DEFAULTPLAYERLIMITPERTABLE uint8 = 5

//================== Dealer ==================\\
//How many decks in a shoe
const DEFAULTDECKPERSHOE = 1

//================== Player ==================\\
//how much money player starts with
const DEFAULTPLAYERSTARTINGCASH float64 = 1000
const DEFAULTPLAYERBET float64 = 10

//3 strikes and you are out
const DEFAULTSTRIKEOUT uint8 = 3

//valid actions
var PlayerAction = map[string]string{
	"hit":      "hit",
	"stand":    "stand",
	"double":   "double",
	"split":    "split",
	"splitAll": "splitAll",
}

var dealerAction = map[string]string{
	"dealSelf": "dealSelf",
	"stand":    "stand",
}
