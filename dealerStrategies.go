//similar to player strategies class
//To use: first init a Dealer's strategy in init function
//call simulate(), which would in turn call one of these strategy functions
package wonging

import (
	"fmt"
)

// type DealerStrategies interface {
// 	standsOnAll17(d *Dealer) *Dealer
// 	hitOnSoft17(d *Dealer) *Dealer
// }
//
type DealerStrategy func(*Dealer) []string

func randomDealerStrategy() (DealerStrategy, string) {
	strategies := []DealerStrategy{standsOnAll17, hitOnSoft17}
	description := []string{"Stands On All 17", "Hit On Soft 17"}
	randomInt := randInt(0, len(strategies))
	return strategies[randomInt], description[randomInt]
}

//Dealer strategies
func standsOnAll17(d *Dealer) []string {
	fmt.Print("[strategy: standsOnAll17]: ")
	if value, _ := d.calculateHandValue(); value < 17 {
		return []string{"dealSelf"}
	}
	return []string{"stand"}
}
func hitOnSoft17(d *Dealer) []string {
	fmt.Print("[strategy: hitOnSoft17]: ")

	if value, soft := d.calculateHandValue(); (value >= 17 && soft) || value < 17 {
		return []string{"dealSelf"}
	}
	return []string{"stand"}
}
