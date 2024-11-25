package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{1, 4, 45, 5, 6, 6, 4, 55, 3, 23, 4, 2, 35}
	sort.Ints(i)
	fmt.Println(i)
}
