package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open a file (this could be any file, just an example)
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Defer the closing of the file until main() completes
	defer file.Close()

	// Perform file operations (this is just a placeholder)
	fmt.Println("File opened successfully!")

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content of files are :", string(content))

	// No need to manually close the file, defer ensures it's closed when the function exits
}
