package main

import "testing"

func BenchmarkMain(b *testing.B) {
	grades := []int32{73, 67, 38, 33}
	for i := 0; i < b.N; i++ {
		gradingStudents(grades)
	}
}
