package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected deck length of 24, but got %d", len(d))
	}

	if d[0] != "King of Spades" {
		t.Errorf("Expected 'King of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs', but got '%v'", d[len(d)-1])
	}
}
