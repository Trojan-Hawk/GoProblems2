// Student Name: Timothy Cassidy
// Student Number: G00333333

// source
// https://golang.org/doc/articles/wiki/

// Used seperate files for the first two parts of the problem
// this will be the main file for the remainder

package main

//To use the net/http package, it must be imported
import (
	//"fmt"
	"net/http"
)

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func handler(w http.ResponseWriter, r *http.Request) {
	// root page
	http.ServeFile(w, r, "index.html")
}

func main() {
	//call handler function
	http.HandleFunc("/", handler)
	http.HandleFunc("/guess", handler)
	http.ListenAndServe(":8080", nil)
}
