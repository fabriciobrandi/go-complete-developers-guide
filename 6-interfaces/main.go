package main

import "fmt"

//In Go you declare interfaces based on the return type of the function
// It means you don't need to especify that englishBot implements the bot interface, it will be implicit
//that you can pass as a receiver to a func called getGreeting that returns a string in this example

//Delaring an interface
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	//In the print greeting you can use any type of receiver to send to the function who expected a bot as a parameter

	printGreeting(eb)
	printGreeting(sb)
}

//Setting custom  implementations to different receivers
func (englishBot) getGreeting() string {
	//VEry custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	//VEry custom logic for generating an english greeting
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
