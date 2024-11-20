package main

import (
	"math"
)

type callbackFunc func(int) int

// cbF used to pass on process
func cbF(i int) int {
	return int(math.Pow(float64(i), float64(i)))
}

// process used to test callback
func process(cb callbackFunc, i int) int {
	return cb(i) + i
}

// iterate in call
func iterate(nums *[]int, cb callbackFunc) {
	for i := range *nums {
		(*nums)[i] = cb(i + 1)
	}
}
