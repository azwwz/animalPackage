package main

import (
	"fmt"
)

// code that use annoymous
// func main() {
// 	c := make(chan int)
// 	go func(){
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)
// }

// code use buffer

func main() {
	c := make(chan int, 0)
	c <- 42
	fmt.Println(<-c)
}
