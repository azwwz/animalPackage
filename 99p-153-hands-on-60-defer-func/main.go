package main

import "fmt"

func main() {
	f1()           //f1
	defer f3(f2()) // f2 -> f3
	f4()           //f4
	defer f5()
	defer f6()
}

func f1() {
	fmt.Println("f1")
}
func f2() int {
	fmt.Println("f2")
	return 1
}
func f3(int) {
	fmt.Println("f3")
}
func f4() {
	fmt.Println("f4")
}
func f5() {
	fmt.Println("f5")
}
func f6() {
	fmt.Println("f6")
}

/*
Hands-on exercise #60 - defer func
● “defer” multiple functions in main
	○ show that a deferred func runs after the func containing it exits.
	○ determine the order in which the multiple defer funcs run
https://go.dev/play/p/tbvSX8qy6TT
curriculum item # 153-hands-on-60-defer-func

*/
