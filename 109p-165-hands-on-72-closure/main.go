package main

import (
	"math"
)

func closure() func() float64 {
	i := 2.0
	return func() float64 {
		i = math.Pow(i, 2)
		return i
	}
}
