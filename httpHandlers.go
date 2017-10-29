package main

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	MessageID string
	Text      string
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	income := decodeIncomingMessage(r)
	outgoing := buildOutgoingMessage(income)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outgoing)
}

// ------------------------------ helpers functions ----------------------------------

func decodeIncomingMessage(r *http.Request) Message {
	decoder := json.NewDecoder(r.Body)
	var income Message
	incomeErr := decoder.Decode(&income)

	if incomeErr != nil {
		panic(incomeErr)
	}

	defer r.Body.Close()
	return Message{income.MessageID, income.Text}
}

func buildOutgoingMessage(m Message) []byte {
	botResponse := parseMessageAndGetContext(m.Text)
	data := Message{"13291231", botResponse}
	jData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jData
}
