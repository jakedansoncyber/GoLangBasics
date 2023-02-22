package main

import (
	"fmt"
	"os"
)

func main() {

	// Try to open file ./test1.txt (this is not the best way to open a file)
	// os.Open returns two values, a *File, and error. We don't need the *File (right now)
	// so we wil ignore it with the _
	_, err := os.Open("./test1.txt")

	// Print error to screen
	// We would use a try catch, but this is for an example
	// nil is good (zero), anything else is bad
	if err != nil {
		fmt.Println("This is the error code:", err)
	}
}
