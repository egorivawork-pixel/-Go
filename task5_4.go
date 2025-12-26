package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 7}

	shapes := []Shape{r, c}

	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}
