// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func countBytes(filePath string) {
}

func main() {
	// Define flags
	filePathCountBytes := flag.String("c", "", "")
	countLines := flag.String("l", "", "")
	countWords := flag.String("w", "", "")

	// Parse the command-line flags
	flag.Parse()

	// Check if -c flag was set
	if *filePathCountBytes == "" && *countLines == "" && *countWords == "" {
		fmt.Println("Usage: ")
		fmt.Println("  -c <file_path> : Get file size in bytes")
		fmt.Println("  -l <file_path> : Count number of lines in the file")
		fmt.Println("  -w <filw_path> : Count number of words in a file")
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
	if *countLines != "" || *countWords != "" {
		var filePath string
		if *countLines != "" {
			filePath = *countLines
		} else {
			filePath = *countWords
		}

		file, err := os.Open(*&filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		lineCount, wordCount := 0, 0
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			lineCount++
			words := strings.Fields(line) // Splits the line into words
			wordCount += len(words)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		// Print results based on flags
		if *countLines != "" {
			fmt.Println(lineCount)
		}
		if *countWords != "" {
			fmt.Println(wordCount)
		}
	}
}
