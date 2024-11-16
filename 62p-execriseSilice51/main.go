package main

import "fmt"

func main() {
	mapt := make(map[string][]string)
	mapt[`bond_james`] = append(mapt[`bond_james`], `shaken, not stirred`, `martinis`, `fast cars`)
	mapt["moneypenny_jenny"] = append(mapt["moneypenny_jenny"], `intelligence`, `literature`, `computer science`)
	mapt["no_dr"] = append(mapt["no_dr"], `cats`, `ice cream`, `sunsets`)
	mapt["fleming_ian"] = append(mapt["fleming_ian"], `steaks`, `cigars`, `espionage`)
	delete(mapt, "fleming_ian")
	for key, val := range mapt {
		for i, valSlice := range val {
			fmt.Printf("%v'%v--%v\n", key, i, valSlice)
		}
		fmt.Println()
	}
}
