package wonging

import (
	"fmt"
)

type Request struct {
	entityType string //possible value "player", "dealer"
	id         uint8  //id of entity
	action     string //possible value "hit", "stand", "double"
	handIndex  uint8  //which hand should we apply the action, 0+
}

func (r *Request) printRequest() {
	fmt.Printf("%s %d requests %s [%d]\n", r.entityType, r.id, r.action, r.handIndex)
}
