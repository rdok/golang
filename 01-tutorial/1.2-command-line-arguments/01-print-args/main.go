package main

import (
	"fmt"
	"os"
	"strings"
)

// $ go run 01-tutorial/01-print-args.go arg1
func main() {
	command := os.Args[0]
	fmt.Println("Command: " + command)

	args := strings.Join(os.Args[1:], ", ")
	fmt.Println("Args: " + args)
}
