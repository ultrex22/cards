package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Length test
	if len(d) != 24 {
		t.Errorf("Expected deck length of 24 but got %d", len(d))
	}
}

func TestNewDeckCard(t *testing.T) {
	d := newDeck()

	//First card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}
	//Last card
	if d[len(d)-1] != "Eight of Diamonds" {
		t.Errorf("Expected Eight of Diamonds but got %v", d[len(d)-1])
	}
}

func TestDeckSaveAndLoad(t *testing.T) {
	_ = os.Remove("_deckTesting")
	d := newDeck()

	err := d.saveToFile("_deckTesting")
	if err != nil {
		t.Errorf("Could not save to file: %v", err)
	}

	loadedFile, err := NewDeckFromFile("_deckTesting")
	if err != nil {
		t.Errorf("Could not open file: %v", err)
	}

	if len(loadedFile) != 24 {
		t.Errorf("Expected loaded file to have a length of 24 but got %d", len(loadedFile))
	}

	_ = os.Remove("_deckTesting")
}
