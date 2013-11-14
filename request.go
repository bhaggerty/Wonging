package wonging

type Request struct {
	entityType string //possible value "player", "dealer"
	id         uint8  //id of entity
	action     string //possible value "hit", "stand", "double"
	handIndex  uint8  //which hand should we apply the action, 0+
}
