package main

import (
	"fmt"
)

func main() {
	cards := deck{"ace of diamonds", newCard()}
	cards = append(cards, "six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "ficarde of diamonds"
}
