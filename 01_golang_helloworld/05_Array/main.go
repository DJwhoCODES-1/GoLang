package main

import "fmt"

func main() {
	fmt.Println("Arrays : Fixed size")

	// var name [5]string
	// name[0] = "Devanshu"
	// name[1] = "Aakash"
	// name[2] = "Ankit"
	// name[3] = "Naman"
	// name[4] = "Shreyansh"
	// fmt.Println("Names:", name)

	// var numbers = [5]int{1, 2, 3}
	// fmt.Println("Numbers:", numbers)
	// fmt.Println("Length of Numbers Array is:", len(numbers))

	var fruits [4]string
	fmt.Println(fruits)
	fmt.Printf("%q\n", fruits) // quoted strings

	var scores [4]int
	fmt.Println(scores)

	var heights [4]float32
	fmt.Println(heights)

	var isDeleted [4]bool
	fmt.Println(isDeleted)
}
