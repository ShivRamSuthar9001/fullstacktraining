package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//greeting
	return "Hola!"
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
