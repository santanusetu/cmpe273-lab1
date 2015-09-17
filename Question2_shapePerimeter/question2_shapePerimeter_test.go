package main

import "testing"
import "fmt"


func test(t *testing.T, one float64, two float64) {
	if one != two {
		t.Error("Expected answer is", one, " but got ", two)
	}
}

func calculate(s Shape) float64 {
	return s.perimeter()
}

func TestShapes(t *testing.T) {
	c := Circle{10}
	r := Rectangle{100, 100}

	test(t, calculate(c), 62.83185307179586)
	test(t, calculate(r), 400)

	c1 := Circle{100}
	fmt.Println(" Perimeter of the Circle is = ", c1.perimeter())
	r1 := Rectangle{200, 100}
	fmt.Println("Perimeter of the Rectangle is = ", r1.perimeter())
}
