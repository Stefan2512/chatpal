package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	print("Listening on port 4419...")
	print(parseMessageAndGetContext("I need some water"))
	http.ListenAndServe(":4419", nil)
}
