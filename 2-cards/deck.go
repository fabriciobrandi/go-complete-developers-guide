package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//This is the sintax of the for loop, because in Go you always have to use the variable declared, if you don't need this variable you can use "_"
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//The receiver is like the satic function in C#, you have do pass to the argument  but not in the parameters
//In this case, any variable of type "deck" now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// -----------------------------This is the returning types of the func, that could be more than one
func deal(d deck, handSize int) (deck, deck) {
	//Deck[startIndexIncluding : upToNotIncluding]
	return d[:handSize], d[handSize:]
}

//Take a slice and converto to string separate with commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Receive a deck and save to file using the package ioutil
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//In go you have to manually create a source to the function rand to generate random values
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
