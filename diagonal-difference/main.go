package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {

	var leftDiagonal int32
	var rightDiagonal int32

	for i := 0; i < len(arr); i++ {
		leftDiagonal += arr[i][i]
		rightDiagonal += arr[i][len(arr)-i-1]
	}
	res := math.Abs(float64(leftDiagonal - rightDiagonal))

	return int32(res)

}

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	// cara kerja
	/*
		11 + 5 + -12 = 4
		4 + 5 + 10  = 19
		4 - 19 = 15

	*/
	fmt.Println(diagonalDifference(arr))
}
