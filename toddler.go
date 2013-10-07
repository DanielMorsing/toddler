//Copyright 2013, Daniel Morsing
//For licensing information, See the LICENSE file

// toddler simulates a inquisitive 5 year old.
//
// While annoying, toddler is very good for questioning your own implicit beliefs while debugging.
// Think of the toddler as a more proactive rubber duck.
//
// Credit goes to David R. MacIver for the original idea . I just put it in a terminal so you can shell out to it.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var initialLines = [...]string{
	"Watcha doing?",
	"What's going on?",
	"What's wrong?",
}

var questions []string

func init() {
	for i := 0; i < 4; i++ {
		questions = append(questions, "Why?")
	}
	for i := 0; i < 2; i++ {
		questions = append(questions, "What does that mean?")
	}
	questions = append(questions, "I don't understand")
}

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	i := rand.Intn(len(initialLines))
	fmt.Println(initialLines[i])
	fmt.Printf("> ")

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		i := rand.Intn(len(questions))
		fmt.Println(questions[i])
		fmt.Printf("> ")
	}
}
