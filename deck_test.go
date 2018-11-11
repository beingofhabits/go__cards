package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck len of 20, but got ", string(len(d)))
	}
}
