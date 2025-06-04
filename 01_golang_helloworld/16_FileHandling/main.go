package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling")

	// Step 1: Create and write to file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file!")
		return
	}
	content := "Hello world by DJ\n"
	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error while writing to file")
		return
	}
	file.Close() // Close after writing
	fmt.Println("Successfully wrote to the file")

	// Step 2: Open file for reading
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file for reading")
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024) // Buffer to hold binary data

	fmt.Println("Reading file contents:")
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file")
			return
		}
		fmt.Print(string(buffer[:n]))
	}
}
