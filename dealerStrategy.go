package wonging

type DealerStrategies interface {
	standsOnAll17(d *Dealer) *Dealer
	hitOnSoft17(d *Dealer) *Dealer
}

//Dealer will stand as long as the card value reaches 17
func standsOnAll17(d *Dealer) {

}

//Dealer will treat the A as 1 and keep hitting
func hitOnSoft17(d *Dealer) {

}
