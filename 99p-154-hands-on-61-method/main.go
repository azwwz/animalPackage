package main

import "log"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	log.Printf("\n name : %v \n age :%v", p.first, p.age)
}

func main() {
	p1 := person{
		first: "todd",
		age:   8,
	}
	p1.speak()
}

/*
Hands-on exercise #61 - method
● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
https://go.dev/play/p/KqOV32q1aC0
curriculum item # 154-hands-on-61-method
*/
