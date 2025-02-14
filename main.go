package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handlers
// Home site
func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

// To view a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// A form to create a snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// To save a new snippet
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Send a 201 status code
	w.WriteHeader(http.StatusCreated)
	// Response body
    w.Write([]byte("Save a new snippet..."))
}

func main() {
	// Initialize a router function
	mux := http.NewServeMux()
	// Define routes
	mux.HandleFunc("GET /{$}", home)                      // $ = Restrict this route to exact matches on / only
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // {} = wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// Start a new web server: TCP network address to listen on and the servemux
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
