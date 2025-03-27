package main

import "fmt"

type bot interface {
	getGreeting() string
}

type (
	englishBot struct{}
	spanishBot struct{}
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// ______________________________________________________________________
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// ______________________________________________________________________
func (eb englishBot) getGreeting() string {
	return "Hello"
}

// ______________________________________________________________________
func (sb spanishBot) getGreeting() string {
	return "Hola"
}
