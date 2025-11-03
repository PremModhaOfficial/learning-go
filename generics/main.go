package main

import (
	"fmt"
	"sync"
)

func main() {
}

func GenericFunction[T any](unknownType T) {
	fmt.Println(unknownType, "is of type")

	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("LOL")
	})
}
