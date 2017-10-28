package main

import (
	"regexp"
	"fmt"
)

// possible statements
var ps = map[string][]string{
	"I need ?(.*)": {
		"Why do you need %s?",
		"Do you really need {0}?",
	},
	"fallbackStatements": {
		"Don't know what to say about this",
	},
}

func parseMessageAndGetContext(s string) string {
	for k := range ps {
		re := regexp.MustCompile(k)
		findings := re.FindStringSubmatch(s)
		if findings != nil {
			interpolatedAnswer := fmt.Sprintf(ps[k][0], findings[1])
			return interpolatedAnswer
		}
	}
	// if no other keys are parsed, fallback to some default statements
	return ps["fallbackStatements"][0]
}

func randomizeMessage(context map[string]string) string {
	// go lang implementation iterates randomly through the keys
	// TODO: should enhance this
	for k := range context {
		return context[k]
	}
	return "nil"
}
