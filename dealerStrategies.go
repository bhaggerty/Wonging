package wonging

type DealerStrategies interface {
	standsOnAll17(d *Dealer) *Dealer
	hitOnSoft17(d *Dealer) *Dealer
}

//Dealer strategies
func standsOnAll17(d *Dealer) {
	if value, _ := d.calculateHandValue(); value < 17 {
		d.dealSelf()
	}
}
func hitOnSoft17(d *Dealer) {
	if value, soft := d.calculateHandValue(); (value >= 17 && soft) || value < 17 {
		d.dealSelf()
	}
}
