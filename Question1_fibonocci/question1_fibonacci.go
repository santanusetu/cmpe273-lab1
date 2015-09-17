package main

import "fmt"

func main() {
	num := 6
	fmt.Println(fibonacci(num))

}

func fibonacci(num int) int{

	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return (fibonacci(num - 1) + fibonacci(num - 2))
	}

}