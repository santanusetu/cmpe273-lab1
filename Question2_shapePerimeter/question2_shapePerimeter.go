package main

import "fmt"
import "math"

func main() {

	rect := Rectangle{50,100}
	circ := Circle{4}
	
	fmt.Println("Rectangle Perimeter =", getPerimeter(rect))
	fmt.Println("Circle Perimeter =", getPerimeter(circ))

}

type Shape interface {
	perimeter() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) perimeter() float64{
	return 2 * ((r.height + r.width))
}

func (c Circle) perimeter() float64{
	return ((2 * math.Pi) * (c.radius))
}

func getPerimeter(shape Shape) float64{
	return shape.perimeter()
}