package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wMute sync.Mutex

func main() {
	count := 0
	const gs = 100
	var wg sync.WaitGroup
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goruntines", runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			wMute.Lock()
			defer wg.Done()
			v := count
			v++
			// runtime.Gosched()
			time.Sleep(time.Microsecond)
			count = v
			wMute.Unlock()
		}()
		fmt.Println("Goruntines", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count: ", count)
}
