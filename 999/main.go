package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type SortInts []int
type SortStrings []string

func (s SortInts) Len() int {
	return len(s)
}
func (s SortInts) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}
func (s SortInts) Less(a, b int) bool {
	return s[a] < s[b]
}

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Sort(SortInts(xi))
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	slices.SortFunc(xs, func(a, b string) int {
		return strings.Compare(a, b)
	})
	fmt.Println(xs)
}
