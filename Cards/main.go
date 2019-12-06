package main

func main() {
	// cards := newDeck()

	// hands, remainingHand := deal(cards, 5)

	// hands.print()
	// remainingHand.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_deck")

	cards := newDeckFromFile("my_deck")
	cards.print()
}
