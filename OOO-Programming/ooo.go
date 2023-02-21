package main

import "fmt"

func main() {

}


type Talk struct {
	//Greetinger
	Greeting string
	Farewell string
}

type Personer interface {
	Hello()
	Bye()
}

func Hello() {
	fmt.Println("Hello")
}

func Bye() {
	fmt.Println("Bye")
}
