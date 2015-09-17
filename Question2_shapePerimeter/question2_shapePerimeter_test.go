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
	
	r2 := Rectangle{50, 50}
	c2 := Circle{5}

	test(t, calculate(r2), 200)
	test(t, calculate(c2), 31.41592653589793)
	
	r3 := Rectangle{20, 20}
	c3 := Circle{4}

	test(t, calculate(r3), 80)
	test(t, calculate(c3), 25.132741228718345)
	
	r1 := Rectangle{200, 100}
	fmt.Println("Perimeter of the Rectangle is = ", r1.perimeter())
	c1 := Circle{100}
	fmt.Println(" Perimeter of the Circle is = ", c1.perimeter())
}
