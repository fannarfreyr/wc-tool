// main.go
package main

import (
	"flag"
	"fmt"
	"os"
)

func countBytes(filePath string) {
}

func main() {
	// Define the -c flag
	filePathCountBytes := flag.String("c", "", "")

	// Parse the command-line flags
	flag.Parse()

	// Check if -c flag was set
	if *filePathCountBytes == "" {
		fmt.Println("No file specified. Use -c <file_path>")
		os.Exit(1)
	}

	// Get file information
	fileInfo, err := os.Stat(*filePathCountBytes)
	if err != nil {
		fmt.Println("Error accessing file:", err)
		os.Exit(1)
	}

	// Print the file size in bytes
	fmt.Println(fileInfo.Size())
}
