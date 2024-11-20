package main

import "testing"

func TestFuncDuration(t *testing.T) {
	FuncDuration(MyFunction)
}

func TestLogger(t *testing.T) {
	Logger(MyFunction)
}

func TestProcessFunc(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	ProcessData(data, callbackFunc)
}
