package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	const gs = 100
	var wg sync.WaitGroup
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goruntines", runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := count
			v++
			runtime.Gosched()
			// time.Sleep(time.Microsecond)
			count = v
			wg.Done()
		}()
		fmt.Println("Goruntines", i, runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("count: ", count)
}
