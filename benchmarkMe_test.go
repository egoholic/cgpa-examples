package main

import (
	"testing"
)


// Run benchmarks with:
//   $ go test -bench=. benchmarkMe.go benchmarkMe_test.go

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}

func benchmarkfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3(n)
	}
	result = r
}

func Benchmark30Fibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}

func Benchmark30Fibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}

func Benchmark30Fibo3(b *testing.B) {
	benchmarkfibo3(b, 30)
}

func Benchmark50Fibo1(b *testing.B) {
	benchmarkfibo1(b, 50)
}

func Benchmark50Fibo2(b *testing.B) {
	benchmarkfibo2(b, 50)
}

func Benchmark50Fibo3(b *testing.B) {
	benchmarkfibo3(b, 50)
}
