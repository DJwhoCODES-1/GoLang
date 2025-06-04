package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Data Conversion")

	var num int = 42
	fmt.Printf("Number %d is of type %T\n", num, num)

	var num1 float64 = float64(num)
	fmt.Printf("Number %f is of type %T\n", num1, num1)

	str1 := strconv.Itoa(num)
	fmt.Printf("Number %s is of type %T\n", str1, str1)

	num2, _ := strconv.Atoi(str1)
	fmt.Printf("Number %d is of type %T\n", num2, num2)
}
