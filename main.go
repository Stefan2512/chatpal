package main

import (
	"net/http"
)

func main() {
	// serve static files
	http.Handle("/", http.FileServer(http.Dir("static")))
	// endpoint for message requests
	http.HandleFunc("/message", MessageHandler)
	print("Listening on port 4419...")
	http.ListenAndServe(":4419", nil)
}
