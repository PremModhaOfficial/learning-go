package main

const filePath = "my_cards.txt"

func main() {
	cards := newDeck()

	cards.suffle()
	cards.saveToFile(filePath)

	newFileDeck := newDeckFromFile(filePath)

	newFileDeck.print()
}
