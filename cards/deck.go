package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) suffle() {
	for i := range d {
		newPos := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return deck(strings.Split(string(bytes), ","))
}

func newDeck() deck {
	cards := deck{}
	cardSutes := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four"}

	for _, sute := range cardSutes {
		for _, value := range cardvalues {
			cards = append(cards, value+" of "+sute)
		}
	}

	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0o666)
}
