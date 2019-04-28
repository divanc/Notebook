package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {

	rect := Rectangle{10, 33}
	circle := Circle{6, 31, 12}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rect))
}
