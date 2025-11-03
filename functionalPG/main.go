package main

import (
	"fmt"
)

func main() {
	fmt.Println(" appplyForAll('prem', 'a', 23) ")
	appplyForAll("prem", "a", 23)
	// appplyForAllTyped("prem", "a", 23) // this will be a error
	appplyForAllTyped("prem", "a", fmt.Sprint(23))
}

// LEARED: the generics help me constraint the relations between types
func appplyForAll(someData ...any) any {
	// this will actulay return diffrent type thab the one reseaved
	for a, i := range someData {
		fmt.Println(a, i)
	}
	return ""
}

func appplyForAllTyped[T any](someData ...T) {
	for a, i := range someData {
		fmt.Println(a, i)
	}
}
