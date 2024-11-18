package main

import "fmt"

type Person struct {
	firstName    string
	lastName     string
	flavorsCream []string
}
type IceCream struct {
	flavors string
}

func main() {
	personOne := Person{
		firstName:    "youyou",
		lastName:     "tang",
		flavorsCream: []string{"herb"},
	}
	personTwo := Person{
		firstName:    "ziqiao",
		lastName:     "lv",
		flavorsCream: []string{"chocolate"},
	}
	fmt.Printf("name : %v %v,\t flavort ice creams : %v\n", personOne.lastName, personOne.firstName, personOne.flavorsCream)
	fmt.Printf("name : %v %v,\t flavort ice creams : %v\n", personTwo.lastName, personTwo.firstName, personTwo.flavorsCream)

}
