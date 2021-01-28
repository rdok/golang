// Prints non unique text of each line, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// $ cat input.txt | go run main.go
func main() {
	duplicateLines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		duplicateLines[input.Text()]++
	}

	fmt.Printf("%s\t%s\n", "Line Number", "Content")
	for line, occurrences := range duplicateLines {
		if occurrences > 1 {
			fmt.Printf("%d\t\t%s\n", occurrences, line)
		}
	}
}
