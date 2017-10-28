package main

import (
	"regexp"
	"fmt"
	"math/rand"
	"time"
)

// possible statements
var ps = map[string][]string{
	"I need ?(.*)": {
		"Why do you need %s?",
		"Do you really need %s?",
	},
	"fallbackStatements": {
		"Don't know what to say about this",
	},
}


// try to parse with every key in the eliza dictionary,
// if so happens, pick a random predefined message in the dictionary
func parseMessageAndGetContext(s string) string {
	for k := range ps {
		re := regexp.MustCompile(k)
		findings := re.FindStringSubmatch(s)
		if findings != nil {
			randomMessageIndex := randomizeMessage(len(ps[k]))
			interpolatedAnswer := fmt.Sprintf(ps[k][randomMessageIndex], findings[1])
			return interpolatedAnswer
		}
	}

	// if no other keys are found, fallback to some default statements
	fallBackKey := "fallbackStatements"
	randomFallbackMessageIndex := randomizeMessage(len(ps[fallBackKey]))
	return ps[fallBackKey][randomFallbackMessageIndex]
}

func randomizeMessage(n int) int {
	// appropriately seed the pseudo-random generator
	// so that at every message to get a different response
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(n)
}
