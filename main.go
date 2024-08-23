package main

import (
	"fmt"
	"log"
)

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Cards in hand:")
	hand.print()

	fmt.Println("\nCards in deck:")
	remainingDeck.print()
	fmt.Println("--------------------------------")

	cards.saveToFile("testSaveFileDeck.txt")
	fmt.Print("Cards saved to file successfully!\n\n")

	newDeck, err := NewDeckFromFile("testSaveFileDeck.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cards opened from file:")

	fmt.Println(newDeck)
	fmt.Println("--------------------------------")
	cards = cards.shuffleDeck()
	cards.print()

}
