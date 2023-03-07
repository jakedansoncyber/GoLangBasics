package main

import (
	"fmt"
)

func main() {

	// length of 3
	// capacity of 3
	// set value when delcaring variable
	courses := []string{
		"First Value",
		"Second Value",
		"Third Value",
	}
	courses = append(courses, "hello")

	// use make function to set length and capacity of slice
	// make accepts 3 values
	// type,
	// length (# of elements),
	// capacity (expected/maximum size) (this is the size of the array underneith the hood)
	// courses := make([]string, 5, 10)
	// set values to courses
	// courses[0] = "First Value"
	// courses[1] = "Second Value"
	// courses[2] = "Third Value"

	// len = length function
	// cap = capacity function
	fmt.Printf("Length of the slice = %d and capacity = %d \n", len(courses), cap(courses))

	for _, i := range courses {
		fmt.Println(i)
	}

}
