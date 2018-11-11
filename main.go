package main

func main() {
	cards := newDeck()

	first, second := deal(cards, 5)

	first.print()
	second.print()
}
