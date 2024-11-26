package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var count int64
	incrementor := 100
	runtime.GOMAXPROCS(1)
	wg.Add(incrementor)
	for i := 0; i < incrementor; i++ {
		go func() {
			defer wg.Done()
			v := count
			v++
			fmt.Println(runtime.GOMAXPROCS(0))
			runtime.Gosched()
			// time.Sleep(time.Millisecond)
			count = v
		}()
	}
	wg.Wait()
	fmt.Println("count", count, runtime.GOMAXPROCS(2))
}
