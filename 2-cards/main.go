package main

//For run this project you need to run both files in the package "go run main.go deck.go"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//Creatting a deck of cards manually
	//var card string = "Ace of Spades"
	//card := newCard()
	//cards = append(cards, "Six of Spades")

	//Save to File
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//New Deck from rexiting file
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	//Deal function
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
