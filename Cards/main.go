package main

func main() {
	cards := newDeck()

	hands, remainingHand := deal(cards, 5)

	hands.print()
	remainingHand.print()

}
