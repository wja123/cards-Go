package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 52 {
    t.Errorf("Expected deck length of 52 but got %v",len(d))
  }

  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card of Ace of Space, but got %v", d[0])
  }

  if d[len(d) - 1] != "King of Diamonds" {
    t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
  }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
 os.Remove("_desktesting")

  d := newDeck()
  d.saveToFile("_decktesting")

  loadedDeck := newDeckFromFile("_decktesting")

 if len(loadedDeck) != 52 {
  t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
 }

  os.Remove("_decktesting")
}
