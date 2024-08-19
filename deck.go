package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Creating a new type of 'deck',
// This will be a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckString := (strings.Join(d, ","))
	return deckString
}

func (d deck) saveToFile(filename string) error {

	// `b` contains everything your file has.
	// This writes it to the Standard Out.
	// os.Stdout.Write(b)

	// You can also write it to a file as a whole.
	err := os.WriteFile(filename, []byte(d.toString()), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func NewDeckFromFile(filename string) (deck, error) {
	deckInBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	deckInString := string(deckInBytes)

	deck := strings.Split(deckInString, ",")
	return deck, err
}
