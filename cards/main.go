package main

import "fmt"

func main() {
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")

	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
