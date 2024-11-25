package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type SortAge []Person

// Len as n
func (s SortAge) Len() int {
	return len(s)
}

// Swap used to swap
func (s SortAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less determine the ascending
// first compara Age Second compara Name
func (s SortAge) Less(i, j int) bool {
	if s[i].Age < s[j].Age {
		return true
	} else if s[i].Age == s[j].Age {
		return s[i].Name < s[j].Name
	}
	return false
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Ames", 32}
	p3 := Person{"Bames", 32}
	p4 := Person{"Cames", 32}
	p5 := Person{"Moneypenny", 27}
	p6 := Person{"Q", 64}
	p7 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4, p5, p6, p7}

	sort.Sort(SortAge(people))
	fmt.Println(people)
}
