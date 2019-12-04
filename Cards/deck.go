package main

import "fmt"

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) returnDeck() deck {
	return d
}
