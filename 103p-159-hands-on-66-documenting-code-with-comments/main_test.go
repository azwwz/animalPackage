package main

import (
	"log"
	"testing"
)

// test
func TestParadise(t *testing.T) {
	got := paradise("Bali")
	want := "My idea of paradise is Bali"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}
