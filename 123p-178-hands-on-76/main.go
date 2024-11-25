package main

import "fmt"

type dog struct {
	first string
}

func (d *dog) walk() {
	fmt.Println("dog can walk", d.first)
}

func (d *dog) run() {
	fmt.Println("dog can run", d.first)
}

type youngin interface {
	walk()
	run()
}

func younginFunc(y youngin) {
	y.run()
	y.walk()
}
func main() {
	d1 := &dog{
		first: "HaHa",
	}
	younginFunc(d1)
}
