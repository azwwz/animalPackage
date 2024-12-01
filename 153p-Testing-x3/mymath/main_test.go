package mymath

import (
	"fmt"
	"testing"
)

type testDate struct {
	date   []int
	expect float64
}

var data []testDate = []testDate{
	{[]int{1, 4, 6, 8, 100}, 6.0},
	{[]int{0, 8, 10, 1000}, 9.0},
	{[]int{9000, 4, 10, 8, 6, 12}, 9.0},
	{[]int{123, 744, 140, 200}, 170.0},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range data {
		n := CenteredAvg(v.date)
		if n != v.expect {
			t.Error("Expect", v.expect, " but got ", n)
		}
	}
}
func ExampleCenteredAvg() {
	for _, v := range data {
		fmt.Println(CenteredAvg(v.date))
	}

	// output:
	// 6
	// 9
	// 9
	// 170
}
func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			CenteredAvg(v.date)
		}
	}
}
