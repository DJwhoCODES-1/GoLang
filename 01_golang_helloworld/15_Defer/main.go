package main

import "fmt"

func main() {
	fmt.Println("Defer : It delays the execution. The method or statement is executed just before the main() is going to end.")

	defer fmt.Println("Starting")
	defer fmt.Println("Middle")
	fmt.Println("Ending")
}

// Defer follows LIFO : Last defer which got added to lifo-stack will be printed first
// It is basically used to close the connections after their use, if we forget to close them in the end
