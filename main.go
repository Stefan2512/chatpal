package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/message", MessageHandler)
	print("Listening on port 4419...")
	http.ListenAndServe(":4419", nil)
}
