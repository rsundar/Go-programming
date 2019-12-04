package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Queen of spades")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
