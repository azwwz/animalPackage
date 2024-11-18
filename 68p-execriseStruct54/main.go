package main

import "fmt"

type Person struct {
	firstName    string
	lastName     string
	flavorsCream []string
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
	for i, values := range personOne.flavorsCream {
		fmt.Printf("%v - %v\n", i, values)
	}

	fmt.Println("----------------")
	Persons := make(map[string]Person)
	Persons[personOne.lastName] = personOne
	Persons[personTwo.lastName] = personTwo
	for key, value := range Persons {
		fmt.Printf("%v -- %v \n", key, value)
		for i, flavors := range value.flavorsCream {
			fmt.Printf("%v - %v\n", i, flavors)
		}
	}

}
