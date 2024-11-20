package main

import (
	"log"
	"testing"
)

func TestClosure(t *testing.T) {
	a := closure()
	a()
	got := a()
	want := 16.0
	if got != want {
		log.Fatalf("got : %v and want : %v", got, want)
	}

}
