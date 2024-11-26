package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("ROUTINEs ", runtime.NumGoroutine())
	fmt.Println("CgoCalls ", runtime.NumCgoCall())
}
