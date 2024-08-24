package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	// "slices"
	"strings"
	"time"
)

// Creating a new type of 'deck',
// This will be a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println("--------------------------------")
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("--------------------------------")
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

func (d deck) shuffleDeck() {
	// making a custom seed/and rand type using time
	source := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(source)

	numOfLoops:=newRand.Intn(10)
	
	// for loop to shuffle the deck randon number of times
	for i:=0; i<numOfLoops; i++ { 
	// for loop to shuffle the deck 
		for loopIndex, card := range d { 
			// if loopIndex == 4 {
			// 	return d
			// }
	
			randIndex := newRand.Intn(len(d) - 1)
			// fmt.Println("randomIndex:", randIndex)
	
			loopCard := card
			// fmt.Println("cardA:", loopCard)
	
			indexCard := d[randIndex]
			// fmt.Println("cardB:", indexCard)
	
			d[loopIndex], d[randIndex] = indexCard, loopCard

			// original way i designed it. works but above is simpler
			// d = slices.Replace(d, randIndex, randIndex+1, loopCard)
			// d = slices.Replace(d, loopIndex, loopIndex+1, indexCard)
		}	
	}
			d.print()

}
