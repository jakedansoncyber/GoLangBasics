package main

import (
	"fmt"
)


// Go doesn't let you create variables outisde of functions (package level)
// So we must declare them in Var() if doing this
var(
	name string
	course string
	module int
	clip int
)
func main() {
	// Grab string from function
	fmt.Println(GetHelloWorld())
	// Have function write string
	WriteHelloWorld()
	// Write string from input
	WriteInput("This is a string that the fucntion writes")

}

// Grabs Hello World String
func GetHelloWorld() string {
	return "Hello World!"
}

// Writes Hello World
func WriteHelloWorld() {
	fmt.Println("Hello World!")
}

// Writes custom input
func WriteInput(input string) {
	fmt.Println(input)
}
