// Student Name: Timothy Cassidy
// Student Number: G00333333

// source
// https://golang.org/doc/articles/wiki/

// Used seperate files for the first two parts of the problem
// this will be the main file for the remainder

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"net/http"
	"html/template"
	"math/rand"
	"strconv"
	"time"
)// imports

// template message struct
type message struct {
	Message string
}// message_struct

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func handler(w http.ResponseWriter, r *http.Request) {
	// root page
	http.ServeFile(w, r, "index.html")
}// handler

// handles all guess.html requests
func guessHandler(w http.ResponseWriter, r *http.Request) {
	// initialise the variable
	randGuess := 0;
	
	// generating the random number
	rand.Seed(int64(time.Now().Nanosecond()))
	
	// If the cookie "target" exists & error handling
	if _, err := r.Cookie("target"); err != nil {
		// if the cookie does not exist generate num between 1 & 20
		randGuess = ((rand.Int() % 19) + 1)

		// convert the generated num to string to store in cookie
		str := strconv.Itoa(randGuess)
		
		
		
		// creating the cookie "target", with a value of num
		cookie := http.Cookie{
			Name:  "target",
			Value: str,
		}// cookie

		// override "target" cookie
		http.SetCookie(w, &cookie)
	}// if
	// Serving the template
	// setting the message struct
	m := message{Message: "Guess a number between 1 and 20: "}
	
	// parsing the guess.tmpl
	tmpl, _ := template.ParseFiles("guess.tmpl")
	
	// execute the template with the message
	tmpl.Execute(w, m)
	
}// guessHandler

func main() {
	// call handler function
	http.HandleFunc("/", handler)
	
	// call the guess handler
	http.HandleFunc("/guess", guessHandler)
	
	// listen for requests on port 8080
	http.ListenAndServe(":8080", nil)
}// main
