package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck len of 16, but got %v", string(len(d)))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected zero value of deck to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last value of deck to be of Four of Clubs, but got %v", d[len(d)-1])
	}
}
