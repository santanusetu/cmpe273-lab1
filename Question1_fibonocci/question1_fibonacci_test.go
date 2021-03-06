package main

import "testing"
import "fmt"

func test(t *testing.T, v int, expectedResult int) {
	res := fibonacci(v)
	if res != expectedResult {
		t.Error("Expected result is ", expectedResult, " - Current result coming as ", res)
	}
}

func TestFibonacci(t *testing.T) {
	test(t, 0, 0)
	test(t, 1, 1)
	test(t, 2, 1)
	test(t, 3, 2)
	test(t, 4, 3)
	test(t, 5, 5)
	test(t, 6, 8)
	test(t, 7, 13)
	test(t, 8, 21)
	test(t, 9, 34)
	test(t, 10, 55)
	test(t, 15, 610)

	fmt.Println(fibonacci(15))
}
