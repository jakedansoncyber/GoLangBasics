package main

import (
	"fmt"
)

func main() {
	// basic for loop
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM!")
			break
		}
		fmt.Println(timer)
		//time.Sleep(1 * time.Second)
	}

	// for range loop
	// unordered slice of string (basically a list)
	coursesInProg := []string{
		"Go",
		"Docker",
		"Docker Networking",
		"Kubernetes"}

	coursesCompleted := []string{
		"Go",
		"Docker"}

	// for range loop
	// basically a foreach loop
	// _ = index value. We want to ingore this, so use an underscore
	// i = data value, so we hang onto this and print out the values inside of coursesInProg
	for _, course := range coursesInProg {
		fmt.Println(course)
		for _, completedCourse := range coursesCompleted {
			if course == completedCourse {
				fmt.Println("Hey we found a clash.", completedCourse, " course has been completed already.")
			}
		}
	}
}
