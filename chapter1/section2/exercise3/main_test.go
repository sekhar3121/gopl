package main

import "testing"

func BenchmarkConcatStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStrings(1000)
	}
}

func BenchmarkBuildStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildStrings(1000)
	}
}
