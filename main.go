package main

func main() {

	card := []string{"Ace of Spades", newcard()}
	card = append(card, "Six of Hearts")

	for i, card := range card {
		println(i, card)
	}

}
