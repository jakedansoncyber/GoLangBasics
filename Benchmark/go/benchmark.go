package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	start := time.Now()
	var myList []int
	for i := 0; i < 100000; i++ {
		myList = append(myList, rand.Intn(2147452437))
	}
	sort.Ints(myList)

	for i := 0; i < 100000; i++ {
		fmt.Println(myList[i])
	}

	duration := time.Since(start)

	fmt.Println(duration)
}
