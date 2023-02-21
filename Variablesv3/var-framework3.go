package main

import (
	"fmt"
)

func main() {
	name := "Jake Danson,"
	course := "Getting started with Go Lang"

	// Pass value by refernce (pointers)
	fmt.Println("\nHi", name, "your current course is:", course)
	updateCourse(&course)
	fmt.Println("You're currently watching", course)
}

func updateCourse(course *string) {
	*course = "Getting started with Docker"
	fmt.Println("Update course to", *course)
	//return *course
}
