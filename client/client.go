package main

import (
	"fmt"
	"github.com/josephyzhou/wonging"
)

func main() {
	deck := new(wonging.Deck).initDeck()
	fmt.Println("This is the client")
}
