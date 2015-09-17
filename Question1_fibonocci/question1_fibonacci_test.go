package main

import "testing"
import "fmt"

func test(t *testing.T, v int, expected int) {
	r := fibonacci(v)
	if r != expected {
		t.Error("Answer should be ", expected, " but it is now ", r)
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
	test(t, 10, 55)

	fmt.Println(fibonacci(15))
}
