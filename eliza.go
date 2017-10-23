package main

import (
	"regexp"
)

// possible statements
var ps = map[string]map[string]string{
	"bored": {
		"somet": "someth",
		"somet2": "someth2",
		"somet3": "someth3",
		"somet4": "someth4",
	},
	"inspired": {
		"da": "nu",
	},
	"expensive": {

	},
	"cheap": {

	},
	"fallbackStatements": {
		"okay": "you must be really bored...",
	},
}

func parseMessageAndGetContext(s string) map[string]string{
	for k := range ps {
		re := regexp.MustCompile(k)
		if re.FindString(s) == k {
			return ps[k]
		}
	}
	// if no other keys are parsed, fallback to some default statements
	return ps["fallbackStatements"]
}


func randomizeMessage(context map[string]string) string{
	// go lang implementation iterates randomly through the keys
	// TODO: should enhance this
	for k := range context {
		return context[k]
	}
	return "nil"
}