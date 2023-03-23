package main

import "fmt"

func main() {
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
