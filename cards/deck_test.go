package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	de := 52
	if len(d) != de {
		t.Errorf("Expected deck length of %d, but got %d", de, len(d))
	}

	fc := "Ace of Spades"
	if d[0] != fc {
		t.Errorf("Expected first card of %v, but got %v", fc, d[0])
	}

	lc := "King of Clubs"
	if d[de-1] != lc {
		t.Errorf("Expected last card of %v, but got %v", lc, d[de-1])
	}

}

func TestSaveToFileAndTestNewDeckFromFile(t *testing.T) {
	de := 52
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != de {
		t.Errorf("Expected %d cards in deck, but got %d", de, len(loadedDeck))
	}
}
