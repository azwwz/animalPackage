package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func main(){
	fmt.Println("CPUs",runtime.NumCPU())
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("GO #1 i want print something out")
	}()
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("GO #2 i want print something out")
	}()
	wg.Wait()
	fmt.Println("going to end")

}