package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for i, v := range m {
		fmt.Printf("key is %v , and value is %v \n", i, v)
	}
}
