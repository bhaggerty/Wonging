//similar to player strategies class
//To use: first init a Dealer's strategy in init function
//call simulate(), which would in turn call one of these strategy functions
package wonging

// type DealerStrategies interface {
// 	standsOnAll17(d *Dealer) *Dealer
// 	hitOnSoft17(d *Dealer) *Dealer
// }
//
type DealerStrategy func(*Dealer) string

func randomDealerStrategy() DealerStrategy {
	strategies := []DealerStrategy{standsOnAll17, hitOnSoft17}
	return strategies[randInt(0, len(strategies)-1)]
}

//Dealer strategies
func standsOnAll17(d *Dealer) string {
	if value, _ := d.calculateHandValue(); value < 17 {
		return "dealSelf"
	}
	return ""
}
func hitOnSoft17(d *Dealer) string {
	if value, soft := d.calculateHandValue(); (value >= 17 && soft) || value < 17 {
		return "dealSelf"
	}
	return ""
}
