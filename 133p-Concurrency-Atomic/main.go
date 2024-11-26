package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// var wMute sync.Mutex

func main() {
	var count int64
	const gs = 100
	var wg sync.WaitGroup
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goruntines", runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {

			defer wg.Done()
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println("count :", atomic.LoadInt64(&count), i)

			// wMute.Unlock()
		}()
		fmt.Println("Goruntines", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count: ", count)
}
