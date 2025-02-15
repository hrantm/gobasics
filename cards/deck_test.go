package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected first card of King of Diamonds but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
