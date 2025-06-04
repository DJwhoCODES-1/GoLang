package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings")

	data := "apple,orange,banana"
	parts := strings.Split(data, ",")

	fmt.Println(parts)

	for index, value := range parts {
		fmt.Println(index, value)
	}

	str1 := "one two three four two tw five"
	count := strings.Count(str1, "tw")
	fmt.Println(count)

	str2 := "   Devanshu Jain     "
	trimmed := strings.TrimSpace(str2)
	fmt.Printf("%s\n", trimmed)

	fName := "Devanshu"
	lName := "Jain"
	result := strings.Join([]string{fName, lName}, " ")
	fmt.Println(result)
}
