// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

type ProgramOptions struct {
	numberOfBytes      bool
	numberOfLines      bool
	numberOfWords      bool
	numberOfCharacters bool
}

type Result struct {
	fileName           string
	numberOfBytes      uint64
	numberOfLines      uint64
	numberOfWords      uint64
	numberOfCharacters uint64
}

var ProgramName = filepath.Base(os.Args[0])

func main() {
	var options ProgramOptions
	numberOfBytes := flag.Bool("c", false, "Used to count number of bytes in the file")
	numberOfLines := flag.Bool("l", false, "Used to count number of lines in the file")
	numberOfWords := flag.Bool("w", false, "Used to count number of words in the file")
	numberOfCharacters := flag.Bool("m", false, "used to count number of characters in the file")

	flag.Parse()
	if *numberOfBytes {
		options.numberOfBytes = true
	}
	if *numberOfLines {
		options.numberOfLines = true
	}
	if *numberOfWords {
		options.numberOfWords = true
	}
	if *numberOfCharacters {
		options.numberOfCharacters = true
	}

	if !options.numberOfBytes && !options.numberOfLines && !options.numberOfCharacters && !options.numberOfWords {
		options.numberOfBytes, options.numberOfLines, options.numberOfWords = true, true, true
	}

	if filename := flag.Arg(0); filename != "" {
		result := processFile(&options, filename)
		fmt.Println(formatOutput(&result, &options))
	} else {
		processStdin(&options)
	}
}

func processStdin(options *ProgramOptions) {
	reader := bufio.NewReader(os.Stdin)
	result := Result{}

	calculate(reader, &result, options)
	fmt.Println(formatOutput(&result, options))
}

func processFile(options *ProgramOptions, filename string) Result {
	result := Result{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	result.fileName = filename

	reader := bufio.NewReader(file)
	calculate(reader, &result, options)
	return result
}

func calculate(fileReader *bufio.Reader, results *Result, options *ProgramOptions) {
	results.numberOfLines = 0
	results.numberOfWords = 0
	results.numberOfBytes = 0
	results.numberOfCharacters = 0

	var prevRune rune

	for {
		runeRead, runeSize, err := fileReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		if options.numberOfBytes {
			results.numberOfBytes += uint64(runeSize)
		}
		if options.numberOfLines && runeRead == '\n' {
			results.numberOfLines++
		}
		if options.numberOfCharacters {
			results.numberOfCharacters++
		}
		if options.numberOfWords {
			if unicode.IsSpace(runeRead) && !unicode.IsSpace(prevRune) {
				results.numberOfWords++
			}
		}
		prevRune = runeRead
	}
	if prevRune != rune(0) && !unicode.IsSpace(prevRune) {
		results.numberOfWords++
	}
}

func formatOutput(results *Result, options *ProgramOptions) string {
	output := ""

	// -l
	if options.numberOfLines {
		output += fmt.Sprintf("%v\t", results.numberOfLines)
	}
	// -w
	if options.numberOfWords {
		output += fmt.Sprintf("%v\t", results.numberOfWords)
	}
	// -c
	if options.numberOfBytes {
		output += fmt.Sprintf("%v\t", results.numberOfBytes)
	}
	// -m
	if options.numberOfCharacters {
		output += fmt.Sprintf("%v\t", results.numberOfCharacters)
	}
	if results.fileName != "" && results.fileName != "-" {
		output += fmt.Sprintf(results.fileName)
	}
	return output
}
