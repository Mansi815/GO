// https:docs.aws.amazon.com/ sdk-for-go/v1/developer-guide/welcome.html
package main

import (
	"fmt"
	"math"
)

// ////  Interface
type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

// define a method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for rectangle (implementation of Shape.area())
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Printf("Circle Area :%.2f\n", getArea(circle))
	fmt.Printf("Rectangle Area :%.2f\n", getArea(rectangle))
}
