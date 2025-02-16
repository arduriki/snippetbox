package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// To serve static files from its directory
	fileServer := http.FileServer(http.Dir("ui/static/"))

	// Register the file server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Other app routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
