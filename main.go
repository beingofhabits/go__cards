package main

func main() {
	cards := newDeck()
	cards.saveToFile("myCards")

	newDeckFromFile("myCards").print()
}
