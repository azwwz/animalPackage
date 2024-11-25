package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "tomson",
		Last:  "xiu",
		Age:   122,
	}
	p2 := Person{
		First: "lakcon",
		Last:  "luis",
		Age:   111,
	}
	pa := []Person{p1, p2}
	jsonPa, err := json.Marshal(pa)
	if err != nil {
		fmt.Println("error = ", err)
	}
	fmt.Println(string(jsonPa))
}
