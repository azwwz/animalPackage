package main

import (
	"log"
	"testing"
)

func TestProcess(t *testing.T) {
	got := process(cbF, 2)
	want := 6
	if got != want {
		log.Fatalf("got : %v and want : %v", got, want)
	}
}

func TestProcess2(t *testing.T) {
	// use annoymous pass in arguments
	a2 := func(a int) int {
		return a + a
	}
	got := process(a2, 2)
	want := 6
	if got != want {
		log.Fatalf("got : %v and want : %v", got, want)
	}
}
func TestIterage(t *testing.T) {
	// use annoymous pass in arguments
	a2 := func(a int) int {
		return a + a
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	iterate(&nums, a2)
	want := []int{2, 4, 6, 8, 10, 12, 14}
	for i, key := range nums {
		if key != want[i] {
			log.Fatalf("got : %v and want : %v", nums, want)
		}
	}

}
