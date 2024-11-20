package main

import "fmt"

func foo() int {
	return 1
}
func bar() (int, string) {
	todd := "mecleod"
	i := 1
	return i, fmt.Sprintf("%s dd", todd)
}
func main() {
	fmt.Printf(string(foo()))
	_, s := bar()
	fmt.Printf(string(s))
}
