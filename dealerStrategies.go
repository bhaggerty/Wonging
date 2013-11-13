package wonging

type DealerStrategies interface {
	standsOnAll17(d *Dealer) *Dealer
	hitOnSoft17(d *Dealer) *Dealer
}

//Dealer strategies
func standsOnAll17(d *Dealer) {
	if d.calculateHandValue() > 17 {

	}
}
func hitOnSoft17(d *Dealer) {

}
