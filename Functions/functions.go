package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "Jake Danson"
	course := "Learning Go"

	fmt.Println(converter(author, course))
}

// input parameters author, course string
// return parameters a, c string
func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}
