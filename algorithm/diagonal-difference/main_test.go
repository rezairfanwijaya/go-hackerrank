package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	actual := digonalDifference(arr)
	assert.Equal(t, int32(2), actual)
}

func BenchmarkAbsolute(b *testing.B) {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	for i := 0; i < b.N; i++ {
		digonalDifference(arr)
	}
}
