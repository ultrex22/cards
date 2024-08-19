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

	// cardBytes := []byte(cards.toString())
	// fmt.Printf("cardBytes: %v\n", cardBytes)

	cards.saveToFile("testSaveFileDeck.txt")
	fmt.Print("Cards saved to file successfully!\n\n")

	nd, err := NewDeckFromFile("testSaveFileDeck.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cards opened from file:")
	fmt.Println(nd)
}
