package main

import (
	"fmt"
	"os"
)

func printErrorAndExit(msg string) {
	err := fmt.Errorf(msg)
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func main() {
	argsLen := len(os.Args)
	if argsLen != 2 {
		printErrorAndExit("Please provide exactly one file")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		printErrorAndExit(err.Error())
	}
	defer file.Close()

	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	printErrorAndExit(err.Error())
	// }
	// fmt.Printf("File info: %#v\n", fileInfo)

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		printErrorAndExit(err.Error())
	}
	// fmt.Printf("File size in bytes: %v\n", len(fileData))

	var lineCount int
	var wordCount int
	var partOfAWord bool
	for _, char := range fileData {
		if char == '\n' {
			lineCount++
			partOfAWord = false
		} else if char != ' ' && !partOfAWord {
			// fmt.Printf("%c ", char)
			wordCount++
			partOfAWord = true
		} else if char == ' ' {
			partOfAWord = false
		}
	}
	fmt.Printf("Number of lines: %v\n", lineCount)
	fmt.Printf("Number of words: %v\n", wordCount)
	fmt.Printf("Number of bytes: %v\n", len(fileData))
}
