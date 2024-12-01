package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		fmt.Println("expect ", 70, " bug got ", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		fmt.Println("expect ", 70, " bug got ", n)
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	//output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
