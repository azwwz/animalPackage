package word

import (
	"fmt"
	"testing"

	"github.com/azwwz/animalPackage/153p-Testing-x2/quote"
)

func TestCount(t *testing.T) {
	n := Count(quote.SunAlso)
	if n != 1349 {
		t.Error("Expect 1349 but got ", n)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// output:
	// 1349
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
