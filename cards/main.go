package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Deck of Cards")
	cards := NewDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	newDeck := newDeckFromFile("my_cards")
	newDeck.Print()
}
