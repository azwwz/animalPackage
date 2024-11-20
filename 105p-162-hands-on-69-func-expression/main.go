package main

import "fmt"

func main() {
	f := func(i int, s string) (int, string) {
		return i, s + "'s (){}()"
	}
	fmt.Println(f(1, "tim"))
}
