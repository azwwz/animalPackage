package main

import "fmt"

func main() {
	x := make([][]string, 2, 50)
	fmt.Printf("len is %v , cap is %v \n", len(x[0]), cap(x[0]))
	x[0] = append(x[0], "James", "Bond", "Shaken, not stirred")
	x[1] = append(x[0], "Miss", "Moneypenny", "I'm 008.")
	for index, arr := range x {
		fmt.Printf("line %v is %v\n\n", index, arr)
		for _, data := range arr {
			fmt.Printf("line %v is %v\n", index, data)
		}
	}
}
