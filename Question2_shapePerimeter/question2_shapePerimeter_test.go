package main

import "testing"
import "fmt"


func test(t *testing.T, expectedResult float64, res float64) {
	if expectedResult != res {
		t.Error("Expected result is", expectedResult, " - but coming as ", res)
	}
}

func calculate(s Shape) float64 {
	return s.perimeter()
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{100, 100}
	c := Circle{10}

	test(t, calculate(r), 400)
	test(t, calculate(c), 62.83185307179586)
	
	r1 := Rectangle{200, 100}
	fmt.Println("Perimeter of the Rectangle is = ", r1.perimeter())
	c1 := Circle{100}
	fmt.Println(" Perimeter of the Circle is = ", c1.perimeter())
}
