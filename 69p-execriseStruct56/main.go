package main

import "fmt"

func main() {
	anonymous := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "liutao",
		friends: map[string]int{
			"good": 12,
		},
		favDrinks: []string{"binghongcha", "jianjiao", "water"},
	}
	fmt.Println(anonymous.first, anonymous.friends["good"], anonymous.favDrinks)
}
