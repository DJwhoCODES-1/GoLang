package main

import (
	"fmt"
)

func modifyValueByReference(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	fmt.Println("Pointers : It is a variable that stores memory address of another variable.")

	// var num int
	// num = 2

	num := 2

	// var ptr *int
	// ptr = &num

	ptr := &num

	fmt.Printf("Num %d is stored at address %x\n", num, ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer does not point to anything")
	}

	var num1 int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&num1)
	fmt.Println(num1)
	modifyValueByReference(&num1)
	fmt.Println(num1)

}
