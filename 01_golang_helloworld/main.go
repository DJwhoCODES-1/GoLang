package main

import (
	"fmt"
	"helloworld/myutil"
)

func main() {
	fmt.Println("Hello Golang World!")

	myutil.PrintMsg("DJwhoCODES")

	var name string = "Devanshu"
	var age int = 24
	var rating float64 = 32.74
	var isDeleted bool = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(rating)
	fmt.Println(isDeleted)

	const fname = "DJwhoCODES"
	// fname = "Jain"  // cannot assign to const variable
	fmt.Println(fname)

	var lname = "Jain"
	lname = "Cena" // can be reassgined
	fmt.Println(lname)

	page := "camera_after_cubicle"
	fmt.Println(page)

	privateMethod()
}

// This function is exported (public) because it starts with a capital letter
func PublicMethod() {
	fmt.Println("This func can be exported")
}

// This function is exported (public) because it starts with a capital letter
func privateMethod() {
	fmt.Println("This func cannot be exported")
}
