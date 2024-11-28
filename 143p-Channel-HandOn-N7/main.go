package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	p := make(chan int)
	go Launches(p)
	for v := range p {
		fmt.Println(" p : ", v)
	}
}
func Launches(p chan<- int) {
	a := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				a++
				p <- a
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			}
		}()
	}
	wg.Wait()
	close(p)

}
