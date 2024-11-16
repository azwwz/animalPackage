package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	if age, ok := m["Jame"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("can't find this one", age)
	}

}
