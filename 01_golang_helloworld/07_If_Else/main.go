package main

import "fmt"

func main() {
	fmt.Println("If Else-If Else")

	z := 10

	if z > 5 {
		fmt.Println("Greater than 5")
	} else if z == 5 {
		fmt.Println("Equal to 5")
	} else {
		fmt.Println("Lesser than 5")
	}
}
