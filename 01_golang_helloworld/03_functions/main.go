package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is simple function")
}

func addNums(a, b int) int { // both a & b must be integers
	return a + b
}

func main() {
	fmt.Println("This iz bizness")

	simpleFunction()

	result := addNums(5, 4)

	fmt.Println(result)
}
