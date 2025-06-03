package main

import "fmt"

func main() {
	age := 24
	name := "Devanshu"
	height := 5.821234

	fmt.Println("age:", age, "|", "name:", name, "|", "height:", height)
	fmt.Println("This is printed in mext line")

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Height is %.2f\n", height)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Name is %s\n", name)

	fmt.Printf("Name is %s, Age is %d, Height is %.4f\n", name, age, height)
}
