package main

import "fmt"

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Queen of spades")

	cards.print()
	fmt.Println(cards.returnDeck())
}

func newCard() string {
	return "Five of diamonds"
}
