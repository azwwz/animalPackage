package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}
func receive(a, b <-chan int) {
	for {
		select {
		case c := <-a:
			fmt.Println("c: ", c)
		case <-b:
			return
		}
	}

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
	}()

	return c
}
