package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("#1 CPUs : ", MaxParallelism())
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 42
		fmt.Println("#2 CPUs : ", runtime.GOMAXPROCS(0))
	}()
	fmt.Println("#3 CPUs : ", MaxParallelism())
	fmt.Println("get num form chan : ", <-c)
	wg.Wait()

}
