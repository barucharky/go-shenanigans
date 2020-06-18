// B''H

/*
go mod init sandbox/magic8ball
go run magic8ball.go
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var answers []string = []string{
	"Very likely",
	"Probably not",
	"I don't see that happening for you",
	"I give it like 50% probability",
	"Unclear. Ask again later",
	"You'll die before that happens",
	"outlook good",
	"I'm tryna think",
}

var badQ map[string]string = map[string]string{
	"what is my name":   "It doesn't matter what your name is",
	"what is your name": "Let's keep this about you",
	"how old am i":      "I think you know that already",
	"am i beautiful":    "You want an honest answer to that?",
	"who will i marry":  "Haha, try to get a date first!",
}

func main() {

	var question string = strings.ToLower(strings.Join(os.Args[1:], " "))
	var isBad bool

	for qMatch, _ := range badQ {
		if question == qMatch {
			isBad = true
		}
	}

	if isBad {
		fmt.Println(badQ[question])
	} else {
		randResponse()
	}

}

func randResponse() {

	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var answer int = rand.Intn(len(answers))

	fmt.Println(answers[answer])
}
