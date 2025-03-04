// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func countBytes(filePath string) {
}

func main() {
	// Define flags
	filePathCountBytes := flag.String("c", "", "")
	countLines := flag.String("l", "", "")

	// Parse the command-line flags
	flag.Parse()

	// Check if -c flag was set
	if *filePathCountBytes == "" && *countLines == "" {
		fmt.Println("Usage: ")
		fmt.Println("  -c <file_path> : Get file size in bytes")
		fmt.Println("  -l <file_path> : Count number of lines in the file")
		os.Exit(1)
	}

	// Handle -c flage (byte count)
	if *filePathCountBytes != "" {
		fileInfo, err := os.Stat(*filePathCountBytes)
		if err != nil {
			fmt.Println("Error accessing file:", err)
			os.Exit(1)
		}
		fmt.Println(fileInfo.Size())
	}

	// Handle -l flag (line count)
	if *countLines != "" {
		file, err := os.Open(*countLines)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		lineCount := 0
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		fmt.Println(lineCount)
	}
}
