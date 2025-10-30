package main

import (
	"fmt"
	"sync"
)

func main() {
	fib := concurrentFib()
	for i := 1; i <= 10000; i++ {
		fmt.Println(i, fib(i))
	}
}

// func fib(number int) int {
// 	baseCase := map[int]int{
// 		1: 1,
// 		2: 1,
// 	}
//
// 	if number < 2 {
// 		return baseCase[number]
// 	}
//
// 	return fib(number-1) + fib(number-2)
// }

type ftype func(int) int

func concurrentFib() ftype {
	sharedCache := map[int]int{
		0: 1,
		1: 1,
		2: 1,
	}

	var wb sync.WaitGroup

	var fib ftype
	fib = func(n int) int {
		wb.Add(1)
		defer wb.Done()
		if n <= 1 {
			return n
		}
		if num, exist := sharedCache[n]; exist {
			return num
		}
		sharedCache[n] = fib(n-1) + fib(n-2)

		return sharedCache[n]
	}

	workerFunc := func(n int) int {
		pipe := make(chan int, 1)
		defer close(pipe)

		wb.Add(1)
		go func() {
			pipe <- fib(n)
		}()
		defer wb.Done()

		return <-pipe
	}

	return workerFunc
}

func fibCached() func(int) int {
	cache := map[int]int{
		1: 1,
	}

	var fib func(int) int

	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if num, exist := cache[n]; exist {
			return num
		}

		cache[n] = fib(n-1) + fib(n-2)

		return cache[n]
	}

	return fib
}
