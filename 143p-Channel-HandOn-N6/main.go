package main

import "fmt"

func main() {
	c := make(chan int)
	putChan(c)
	getChan(c)
}
func putChan(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
func getChan(c <-chan int) {
	for v := range c {
		fmt.Println("c: ", v)
	}
}
