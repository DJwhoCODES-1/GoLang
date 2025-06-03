package main

import "fmt"

func main() {
	fmt.Println(`For Loop : Go does not suport "while" & "do-while"`)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for loop is used as while loop
	// counter := 10
	// for {
	// 	if counter == 30 {
	// 		break
	// 	}
	// 	fmt.Println(counter)
	// 	counter++
	// }

	numbers := []int{100, 200, 300, 400}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// name := "Devanshu Jain"
	// for index, char := range name {
	// 	fmt.Println(index, char)   // this will print ascii
	// }

	name := "Devanshu Jain"
	for index, value := range name {
		fmt.Printf("%d : %c\n", index, value)
	}
}
