package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, value := range x {
		fmt.Printf("the VALUE - %v and the TYPE - %T \n", value, value)
	}
	fmt.Printf("%T \t %#v\n", x, x)
	fmt.Printf("%T \t %v\n", x, x)
}
