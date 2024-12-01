package main

import (
	"fmt"

	"github.com/azwwz/animalPackage/153p-Testing-x2/quote"
	"github.com/azwwz/animalPackage/153p-Testing-x2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
