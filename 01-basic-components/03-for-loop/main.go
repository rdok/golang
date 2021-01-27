package main

import (
	"fmt"
	"strconv"
)

func main() {
	var output, currentIndex, delimiter string
	delimiter = ", "

	for index := 1; index < 13; index++ {
		currentIndex = strconv.Itoa(index)
		output += currentIndex + delimiter
		fmt.Println("Index: " + currentIndex)
	}

	fmt.Println("Output: " + output)
}
