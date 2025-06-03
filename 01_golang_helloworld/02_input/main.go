package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name)

	reader := bufio.NewReader(os.Stdin)
	fullName, _ := reader.ReadString('\n')
	fmt.Println("Full Name is", fullName)
}
