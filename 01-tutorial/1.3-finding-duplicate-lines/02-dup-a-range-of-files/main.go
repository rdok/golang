package main

import (
	"bufio"
	"fmt"
	os "os"
)

func main() {
	duplicateLines := make(map[string]int)
	files := os.Args[1:]

	// $ cat input.txt | go run main.go
	if len(files) == 0 {
		handleStandardInput(os.Stdin, duplicateLines)
	} else {
		// $ go run main.go input.txt
		handleFilesInput(files, duplicateLines)
	}

	printDup(duplicateLines)
}

func printDup(duplicateLines map[string]int) {
	for duplicateLine, lineNumber := range duplicateLines {
		if lineNumber > 1 {
			fmt.Printf("%d\t%s\n", lineNumber, duplicateLine)
		}
	}
}

func handleFilesInput(files []string, lines map[string]int) {
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		handleStandardInput(file, lines)
	}
}

func handleStandardInput(stdin *os.File, lines map[string]int) {
	input := bufio.NewScanner(stdin)
	for input.Scan() {
		lines[input.Text()]++
	}
}
