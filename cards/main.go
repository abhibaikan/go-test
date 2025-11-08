package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Deck of Cards")
	cards := NewDeck()
	cards.Print()
}
