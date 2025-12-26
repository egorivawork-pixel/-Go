package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{radius: 10}

	fmt.Printf("Радиус: %.2f\n", c.radius)
	fmt.Printf("Площадь круга: %.2f\n", c.Area())
}
