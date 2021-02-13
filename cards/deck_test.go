package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(d))
	}

	if d[0] != "1 de Espada" {
		t.Errorf("Expected '1 de Espada', but got %v", d[0])
	}

	if d[len(d)-1] != "12 de Copa" {
		t.Errorf("Expected '12 de Copa', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const FILENAME = "_decktesting"
	os.Remove(FILENAME)

	d := newDeck()
	d.saveToFile(FILENAME)

	loadedDeck := newDeckFromFile(FILENAME)

	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards from file, but got %v", len(loadedDeck))
	}

	os.Remove(FILENAME)
}
