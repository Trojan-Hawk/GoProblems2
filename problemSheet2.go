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
	"html/template"
)

// template message struct
type message struct {
	Message string
}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func handler(w http.ResponseWriter, r *http.Request) {
	// root page
	http.ServeFile(w, r, "index.html")
}

// handles all guess.html requests
func guessHandler(w http.ResponseWriter, r *http.Request) {
	// root page
	// http.ServeFile(w, r, "guess.html")
	
	// setting the nessage struct
	m := message{Message: "Guess a number between 1 and 20: "}
	
	// parsing the guess.tmpl
	tmpl, _ = tmpl.ParseFiles("guess.tmpl")
	
	// execute the template with the message
	tmpl.Execute(w, m)
	
}

func main() {
	// call handler function
	http.HandleFunc("/", handler)
	
	// call the guess handler
	http.HandleFunc("/guess", guessHandler)
	
	http.ListenAndServe(":8080", nil)
}
