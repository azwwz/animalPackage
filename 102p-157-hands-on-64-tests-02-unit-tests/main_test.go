package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(1, 2)
	want := 3
	if got != want {
		log.Fatalf("error - got : %v and want : %v \n", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(4, 2)
	want := 2
	if got != want {
		log.Fatalf("error - got : %v and want : %v \n", got, want)
	}
}
