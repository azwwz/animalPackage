package main

import (
	"fmt"
	"log"
	"time"
)

// FuncDuration wrappe func() conpute the elapse of the func
func FuncDuration(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println("duration : ", elapsed)
}

// Logger wrappe receive func() and add some log before and after func()
func Logger(f func()) {
	log.Printf("Starting execution...")
	f()
	log.Printf("Execution completed...")
}

// ProcessData uesd to handle []int data
func ProcessData(data []int, callback func(int)) {
	for _, value := range data {
		callback(value)
	}
}

// callBackFunc passed to ProcessData to handle int
var callbackFunc = func(num int) {
	fmt.Println("Processing number:", num)
}

// MyFunction used for test
func MyFunction() {
	fmt.Println("Hello GoPher")
	time.Sleep(time.Second * 2)
	fmt.Println("MyFunction end")
}
