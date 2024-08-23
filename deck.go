package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
)

// Creating a new type of 'deck',
// This will be a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

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

		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	deckInString := string(deckInBytes)

	loadedDeck := strings.Split(deckInString, ",")
	return loadedDeck, err
}

func (d deck) shuffleDeck() deck {
	for loopIndex, card := range d {
		randIndex := rand.Intn(len(d) - 1)
		fmt.Println("randomIndex:", randIndex)
		cardA := card
		fmt.Println("cardA:", cardA)
		cardB := d[randIndex]
		fmt.Println("cardB:", cardB)
		d = slices.Replace(d, randIndex, randIndex, cardA)
		d.print()
		d = slices.Replace(d, loopIndex, loopIndex, cardB)
		d.print()
	}
	return d
}
