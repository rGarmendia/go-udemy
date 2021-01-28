package main

import "fmt"

// if you have any TYPE that matches with the METHOD declare in a interface,
// then that TYPE inherit the bot type, and you can abstract the methods
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// You do not have to make a instance of the receiver type if you are not going to use it
func (englishBot) getGreeting() string {
	//VERY custom logic for getting english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Todo fino?"
}
