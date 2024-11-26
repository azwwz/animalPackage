package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var smute sync.Mutex

func main() {
	var count int64
	incrementor := 100
	wg.Add(incrementor)
	for i := 0; i < incrementor; i++ {
		go func() {
			smute.Lock()
			defer wg.Done()
			v := count
			v++
			// runtime.Gosched()
			time.Sleep(time.Millisecond)
			count = v
			smute.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}
