// Student Name: Timothy Cassidy
// Student Number: G00333333

// source
// https://golang.org/doc/articles/wiki/

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"net/http"
)

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Guessing Game! ")
}

func main() {
	//call handler function
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
