package main

import "fmt"

type bot interface {
	getGreeting() string
	// getBotVersion() float64
	// respondToUser() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating greeting
	return "Hello There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
