package main

import (
	//"net/http"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	print("Listening on port 4419...")
	http.ListenAndServe(":4419", nil)
}
