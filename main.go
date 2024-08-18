package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	cardBytes := []byte(cards.toString())
	fmt.Printf("cardBytes: %v\n", cardBytes)
}
