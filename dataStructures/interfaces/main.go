package main

import "fmt"

type (
	englishBot struct{}
	spanishBot struct{}
)

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "englishBot"
}

func (spanishBot) getGreeting() string {
	return "spanishBot"
}

func pringGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	pringGreeting(eb)
	pringGreeting(sb)
}
