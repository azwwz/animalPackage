package main

import "fmt"

func main() {
	mapt := make(map[string][]string)
	/*
		mapt[`bond_james`] = append(mapt[`bond_james`], `shaken, not stirred`, `martinis`, `fast cars`)
		mapt["moneypenny_jenny"] = append(mapt["moneypenny_jenny"], `intelligence`, `literature`, `computer science`)
		mapt["no_dr"] = append(mapt["no_dr"], `cats`, `ice cream`, `sunsets`)
	*/
	mapt[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	mapt["moneypenny_jenny"] = []string{`intelligence`, `literature`, `computer science`}
	mapt["no_dr"] = []string{`cats`, `ice cream`, `sunsets`}

	for key, val := range mapt {
		for i, valSlice := range val {
			fmt.Printf("%v'%v--%v\n", key, i, valSlice)
		}
		fmt.Println()
	}
}
