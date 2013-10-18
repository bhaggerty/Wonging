package main

import (
	"fmt"
	"github.com/josephyzhou/wonging"
)

func main() {
	fmt.Println("This is the client")

	new(wonging.Deck).InitDeck()
}
