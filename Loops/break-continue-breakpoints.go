package main

import (
	"fmt"
	"time"
)

func main() {
	// for loop with continue
	// prints only odd numbers
	for timer := 10; timer >= 0; timer-- {
		// checks if timer is even nnumber
		if timer%2 == 0 {
			continue // skips fmt.println * time.Sleep which only shows odd numbers
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	// for loop with explosion after 10 seconds
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM!")
			break // exits for loop here
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	// for loop with breakpoint set named breakPoint
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	breakPoint:
		for j := 1; j < 10; j++ {
			fmt.Println("j =", j)
			if j%2 == 0 {
				break breakPoint
			}
		}
	}

}
