package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// $ go run main.go input.txt input-beta.txt
func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			id := fmt.Sprintf("%s %s", filename, line)
			counts[id]++
		}
	}

	for line, occurrences := range counts {
		if occurrences > 1 {
			fmt.Printf("%d\t%s\n", occurrences, line)
		}
	}
}
