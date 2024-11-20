package main

import (
	"log"
	"math"
)

type SHAPE interface { // any type has AREA is my type
	AREA() float64
}
type SQUARE struct { // a square
	length float64
	width  float64
}

func (s SQUARE) AREA() float64 { // square has area . also a shape
	a := s.length * s.width
	return a
}

type CIRCLE struct { // a circle
	radius float64
}

func (c CIRCLE) AREA() float64 { // circle has area . also a shape
	a := math.Pow(c.radius, 2) * math.Pi
	return a
}

func INFO(shap SHAPE) { // take any value about shape , shape will only use mathed AREA()
	log.Printf("area is : %f \n", shap.AREA())
}

func main() {

	squ := SQUARE{ // create a squ also a shape
		length: 12.2,
		width:  13.3,
	}
	cir := CIRCLE{ // create a circle als a shape
		radius: 3.3,
	}
	INFO(squ) // pass a SQUARE of shape
	INFO(cir) // pass a CIRCLE of shape

}

/*
Hands-on exercise #62 - interfaces
● create a type SQUARE
○ length float64
○ width float64
● create a type CIRCLE
○ radius float64
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
■ math.Pi
■ math.Pow
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
https://go.dev/play/p/8BFxl2GXgcw
curriculum item # 155-hands-on-62-interfaces

*/
