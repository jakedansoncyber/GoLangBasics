package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	course := "Kubernetes"

	// Switch case statement. Notice there is no break, this is already implied after it hits
	// a case or the default statement  ("No breaking and Fallthrough, implicit break")
	switch course {
	case "kubernetes":
		fmt.Println("Case 1. kubernetes with lowercase \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2. Kubernetes with uppercase \"K\".")
	case "K8s":
		fmt.Println("Case 3. Kubernetes with short name \"Kates\".")
	case "Istio":
		fmt.Println("Case 4. kubernetes with lowercase \"l\".")
		//fallthrough <- if you put a fallthrough before the default, it will run the default
	default:
		fmt.Println("Hit the default")
	}

	// Switch case statement with a "fallthrough"
	// This allows multiple statements to be called
	// It is best practice to not do this and to instead match multiple values in the same case statement.
	switch course {
	case "kubernetes":
		fmt.Println("Case 1. kubernetes with lowercase \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2. Kubernetes with uppercase \"K\".")
		fallthrough // will allow the next statement to run
	case "K8s":
		fmt.Println("Case 3. Kubernetes with short name \"Kates\".") // will stop here unless we add another fallthrough below
	case "Istio":
		fmt.Println("Case 4. kubernetes with lowercase \"l\".")
	default:
		fmt.Println("Hit the default")
	}

	// Best practice switch case statement (idiomatic)

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got the following even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got the following odd number -", tmpNum)
	default:
		fmt.Println("We got a zero")
	}

}

// <returns> A random Int <returns>
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
