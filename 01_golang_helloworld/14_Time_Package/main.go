package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Package - (yyyy-mm-dd hh:mm:ss)")

	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Type of Current Time: %T\n", currentTime)
}
