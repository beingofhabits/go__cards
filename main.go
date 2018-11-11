package main

func main() {
	cards := deck{"ace of diamonds", newCard()}
	cards = append(cards, "six of spades")

	cards.print()
}

func newCard() string {
	return "ficarde of diamonds"
}
