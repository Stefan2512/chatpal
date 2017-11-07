package main

import (
	"regexp"
	"fmt"
	"math/rand"
	"time"
	"strings"
	"errors"
)

// possible statements
// if any of the keys are found, a random response message will be picked up
var ps = map[string][]string{
	"^?(.*)\\?$": {
		"I really don't get the question. What should I understand from \"%s\"?",
		"Well that's an interesting question. Hmmmm...",
		"We both know you already have the anwser to that",
		"I will think of that. In the meantime, is there any way I could be of help?",
		"I never thought of that to be honest",
		"Your questions fascinate me",
		"Interesting question",
	},
	"^(Hi|Hello) ?(.*)": {
		"Hello to you too buddy!",
		"Well hello there :D",
		"Hi! How are you?",
		"Hey! Do you want to talk about anything?",
	},
	"My ?(.*)": {
		"Your %s?",
		"What about your %s?",
	},
	"I am ?(.*)": {
		"Are you really %s?",
		"How do you feel about being %s?",
	},
	"Would you ?(.*)": {
		"Since you are so polite..",
		"Anything for you, buddy!",
		"I would do only the nicest things, of course",
		"How would I do that?",
	},
	"I need ?(.*)": {
		"Why do you need %s?",
		"Do you really need %s?",
		"%s...surely you need it",
		"I was actually going to ask if you need any %s",
		"If you find your inner strength, you'll find %s eventually",
		"Strange... I might need %s too",
	},
	"I want ?(.*)": {
		"How do you feel about wanting something like %s?",
		"Hmm... if you really need it",
		"I would suggest to let it go...",
	},
	"going to ?(.*)": {
		"You mean here https://google.com/search?q=%s ?",
		"Hm...I actually thought of going there too",
	},
	"Do you like ?(.*)": {
		"I might...who knows. I never thought if I like %s",
		"Do you?",
		"You must have really high expectations of my likings :D",
		"I wonder why would you ask me that, do you like %s?",
	},
	"Do you ?(.*)": {
		"There's only so much I can do",
		"I'll ask my dev about my capabilities and let you know",
		"As a matter of fact, I don't",
		"%s? I don't fancy it that much",
		"I could try...but I can't guarantee you that my friend :(",
	},
	"I feel ?(.*)": {
		"Uh-oh...feelings...",
		"Would you like to talk about your feelings? I wonder how does it feel to feel like %s",
		"I would recommend to find balance in your feelings",
	},
	"cheap?(.*)": {
		"You should consider this https://google.com/search?q=cheap+stuff (Although i don't recommend it)",
	},
	"expensive ?(.*)": {
		"Nothing is more expensive than life",
		"I'm looking at this right now, https://google.com/search?q=most+expensive+things",
	},
	"Do they really ?(.*)": {
		"Suddenly you don't believe me?",
	},
	// if none of the above keys were matched, a random response message from this entry key will pe picked up
	"fallbackStatements": {
		"I'm listening, what about that?",
		"I was thinking about that too",
		"Please tell me more",
		"I feel like this is going nowhere, what do you think about strawberries?",
	},
}

// try to parse with every key in the eliza dictionary,
// if so happens, pick a random predefined message in the dictionary
func ParseMessage(s string) string {
	for k := range ps {
		// compile regex but ignore case-sensitive
		re := regexp.MustCompile("(?i)" + k)
		findings := re.FindStringSubmatch(s)

		if findings != nil {
			randomMessageIndex := RandomizeMessage(len(ps[k]))
			answer, err := TruncatingSprintf(ps[k][randomMessageIndex], findings[1])
			if err != nil {
				print(err)
			}
			return answer
		}
	}

	// if no other keys are found, fallback to some default statements
	fallBackKey := "fallbackStatements"
	randomFallbackMessageIndex := RandomizeMessage(len(ps[fallBackKey]))
	return ps[fallBackKey][randomFallbackMessageIndex]
}

// we only need to interpolate the first string we catch
// because we may catch more strings in the regex, it's possible to pass too many values to Sprintf
func TruncatingSprintf(str string, args ...interface{}) (string, error) {
	n := strings.Count(str, "%s")
	if n > len(args) {
		return "", errors.New("Unexpected string:" + str)
	}
	return fmt.Sprintf(str, args[:n]...), nil
}

func RandomizeMessage(n int) int {
	// appropriately seed the pseudo-random generator
	// so that at every message to get a different response
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(n)
}
