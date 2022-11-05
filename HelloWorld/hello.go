// This declares that this is an application and not a package
// Packages are usually called upon by applications as packages aren't
// going to be ran standalone.
package main 

// Package that allows us to print to console
import "fmt"

// Applications entry point
func main(){
	fmt.Println("Hello World!")
}