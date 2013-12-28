package wonging

import (
	"fmt"
)

type Request struct {
	entityType string   //possible value "player", "dealer"
	id         uint8    //id of entity
	action     []string //possible value "hit", "stand", "double", can be for multiple hands
	handIndex  []uint8  //which hand should we apply the action, 0+, , can be for multiple hands
}

func (r *Request) printRequest() {
	fmt.Println(r.description())
}

func (r *Request) description() string {
	return fmt.Sprintf("%s %d requests %s %d", r.entityType, r.id, r.action, r.handIndex)
}
