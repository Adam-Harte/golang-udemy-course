package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of spades" {
		t.Errorf("Expected first card of Ace of spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of clubs" {
		t.Errorf("Expected last card of Four of clubs, but got %v", d[len(d) - 1])
	}
}