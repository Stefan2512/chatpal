package main

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	Text      string
}

// handle ingoing and outgoing messages with this handler
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	income := decodeIncomingMessage(r)
	outgoing := buildOutgoingMessage(income)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outgoing)
}

// ------------------------------ helpers functions ----------------------------------

// decode the message send to the server as JSON
func decodeIncomingMessage(r *http.Request) Message {
	decoder := json.NewDecoder(r.Body)
	var income Message
	incomeErr := decoder.Decode(&income)

	if incomeErr != nil {
		panic(incomeErr)
	}

	defer r.Body.Close()
	return Message{income.Text}
}

// build the message as JSON to be send back to the client
func buildOutgoingMessage(m Message) []byte {
	botResponse := ParseMessage(m.Text)
	data := Message{botResponse}
	jData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jData
}
