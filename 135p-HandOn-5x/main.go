package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var count int64
	incrementor := 100
	wg.Add(incrementor)
	for i := 0; i < incrementor; i++ {
		go func() {

			defer wg.Done()
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println("count : ", atomic.LoadInt64(&count))

		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}
