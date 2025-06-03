package main

import "fmt"

func main() {
	fmt.Println("Slices : Flexible & Dynamic sized arrays")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	fmt.Printf("Type: %T\n", numbers)
	fmt.Println("Length:", len(numbers))

	numbers = append(numbers, 10, 9, 8, 7, 6)
	fmt.Println("Numbers:", numbers)

	names := []string{}
	names = append(names, "Devanshu", "Shreyansh")
	fmt.Println("Names:", names)

	// scores := make([]int, length, capacity)
	scores := make([]int, 3, 5)
	fmt.Println("Scores:", scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	scores = append(scores, 4)
	scores = append(scores, 98)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	scores = append(scores, 938)

	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

}
