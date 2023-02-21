package main

import (
	"fmt"
)

func main() {
	bestFinishes := championshipFinishes(12, 5, 6, 6, 7, 3)

	fmt.Println(bestFinishes)
}

// This is a variadic function that expects ANY number of ints
// It needs 3 periods before the type, they care called elipses
// The input parameter finishes gets stored as a "slice" of ints
// <returns> Best finish (lowest value given) <returns>
func championshipFinishes(finishes ...int) int {
	// Set best to the first int in the slice
	best := finishes[0]

	// Loop through to find best finish
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
