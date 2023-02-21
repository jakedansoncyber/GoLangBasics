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
var ()

func main() {

	name := "Jake Danson"
	course := "Getting started with Go Lang"
	module := "4"
	clip := 2
	//courseComplete := false

	// Grab string from function
	fmt.Println(GetHelloWorld())
	// Have function write string
	WriteHelloWorld()
	// Write string from input
	WriteInput("This is a string that the fucntion writes")
	WriteInput(name + course)
	// Converting modulev2 into a string and storing into module
	// since it hasn't been initialized yet
	iModule, errorCodeCapturedFromAtoi := strconv.Atoi(module)

	// Checking if the error code from the strconv function
	// is equal to nil (nothing or zero)
	if errorCodeCapturedFromAtoi == nil {
		total := iModule + clip
		fmt.Println("Module + clip = ", total)
	}

	fmt.Println("Memory address of the course variable: ", &course)

	// Creating a pointer variable
	var ptr *string = &course

	fmt.Println("The memory of ptr = ", ptr, " The value of ptr = ", *ptr)

	*ptr = "New value"

	fmt.Println("The memory of ptr = ", ptr, " The value of ptr = ", *ptr)


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
