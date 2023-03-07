package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye for now... ")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8080", nil)
}
