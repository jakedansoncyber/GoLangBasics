package main

import (
	"fmt"
	"strconv"
)

// Go doesn't let you create variables outisde of functions (package level)
// So we must declare them in Var() if doing this
var (
	name   string
	course string
	module int
	clip   int
)

// Or we can set variables and the program interprets the type for us
var (
	namev2   = "jake"
	coursev2 = "Go Lang!"
	modulev2 = "10" // put as string for Atoi function below
	clipv2   = 2
)

func main() {
	// Grab string from function
	fmt.Println(GetHelloWorld())
	// Have function write string
	WriteHelloWorld()
	// Write string from input
	WriteInput("This is a string that the function writes")

	// Converting modulev2 into a string and storing into module
	// since it hasn't been initialized yet
	// The Error code variable is used as the strconv.Atoi function
	// returns an error code AND an integer, so you have to have two variables
	// to store the error code and the integer
	module, errorCodeCapturedFromAtoi := strconv.Atoi(modulev2)

	// Checking if the error code from the strconv function
	// is equal to nil (nothing or zero)
	if errorCodeCapturedFromAtoi == nil {
		total := module + clip
		fmt.Println("Module + clip = ", total)
	}

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
