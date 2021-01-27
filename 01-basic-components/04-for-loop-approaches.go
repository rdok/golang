package main

import (
	"fmt"
	"os"
)

// $ go run %
func main() {
	output, separator := "", ", "

	for _, arg := range os.Args[1:] {
		output += arg + separator
	}

	fmt.Println("Command arguments: " + output)
}
