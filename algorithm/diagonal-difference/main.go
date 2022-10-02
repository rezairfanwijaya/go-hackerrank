package main

import (
	"log"
	"math"
)

func digonalDifference(arr [][]int32) int32 {
	var diagonalLeft int32 = 0
	var diagonalRight int32 = 0

	for index := range arr {
		diagonalLeft += arr[index][index]
		diagonalRight += arr[index][len(arr)-index-1]
	}

	tmp := diagonalLeft - diagonalRight
	result := math.Abs(float64(tmp))
	return int32(result)
}

func main() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	log.Println(digonalDifference(arr))
}
