package main

import "fmt"

func divideNums(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("This module is about error handling.")

	result, err := divideNums(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	// result, _ := divideNums(10, 0)
	// fmt.Println(result)
}
