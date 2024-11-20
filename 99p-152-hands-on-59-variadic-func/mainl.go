package main

import "fmt"

func main() {
	a := make([]int, 0, 100)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)
	variadicFunc(a...)
}

func variadicFunc(a ...int) {
	var t int // total

	for i, k := range a {
		fmt.Printf("i = %v --- k = %v\n", i, k)
		t += k
	}
	fmt.Printf("the sum is :  %v ", t)
}
