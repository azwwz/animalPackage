package main

import "fmt"

type Person struct {
	last string
}

func (p *Person) speak() {
	fmt.Println("My lastName: ", p.last)
}

type Human interface {
	speak()
}

func saySometing(h Human) {
	h.speak()
}
func main() {
	// p := Person{"xlen"}
	pp := &Person{"Pxlen"}
	// saySometing(p)
	saySometing(pp)
}
