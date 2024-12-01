package main

import (
	"fmt"

	"github.com/azwwz/animalPackage/153p-Testing-x2/quote"
	"github.com/azwwz/animalPackage/153p-Testing-x2/word"
)

func main() {

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
	fmt.Println(word.Count(quote.SunAlso))

}
