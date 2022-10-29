package main

import "testing"

func BenchmarkSircularArrayRotation(b *testing.B) {
	a := []int32{3, 4, 5, 6}
	k := 1
	queries := []int32{1, 2}

	for i := 0; i < b.N; i++ {
		sircularArrayRotation(a, int32(k), queries)
	}
}
