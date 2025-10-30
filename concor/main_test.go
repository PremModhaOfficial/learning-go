package main

import (
	"testing"
)

func BenchmarkConcurrentFib(b *testing.B) {
	fib := concurrentFib()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}

func BenchmarkFibCached(b *testing.B) {
	fib := fibCached()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}

func BenchmarkConcurrentFibLarge(b *testing.B) {
	fib := concurrentFib()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func BenchmarkFibCachedLarge(b *testing.B) {
	fib := fibCached()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}
