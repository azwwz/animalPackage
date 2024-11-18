package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

func main() {
	vehicle1 := vehicle{
		engine: engine{true},
		make:   "easy",
		model:  "su7",
		doors:  "iron",
		color:  "blue",
	}
	vehicle2 := vehicle{
		engine: engine{false},
		make:   "difficulty",
		model:  "su8",
		doors:  "plastics",
		color:  "black",
	}
	fmt.Println(vehicle1.electric, vehicle1.make, vehicle1.model, vehicle1.doors, vehicle1.color)
	fmt.Println(vehicle2.electric, vehicle2.make, vehicle2.model, vehicle2.doors, vehicle2.color)

}
