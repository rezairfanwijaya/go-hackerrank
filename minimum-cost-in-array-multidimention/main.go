package main

import (
	"fmt"
	"math"
)

func minCost(input [][]int) int {
	// ambil dimensi array
	row := len(input)       // m
	column := len(input[0]) // n

	// membuat array tiruan untuk menampung minimun cost
	minimumCostTable := make([][]int, row) // memo
	for i := range minimumCostTable {
		minimumCostTable[i] = make([]int, column)
	}

	// init minimum cost table dengan nilai params di pojok kiri atas
	minimumCostTable[0][0] = input[0][0]

	// Fill the memoization table
	for i := 1; i < row; i++ {
		minimumCostTable[i][0] = minimumCostTable[i-1][0] + input[i][0]
	}
	for j := 1; j < column; j++ {
		minimumCostTable[0][j] = minimumCostTable[0][j-1] + input[0][j]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			minimumCostTable[i][j] = input[i][j] + int(math.Min(float64(minimumCostTable[i-1][j]), float64(minimumCostTable[i][j-1])))
		}
	}

	return minimumCostTable[row-1][column-1]
}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	minCost := minCost(input)
	fmt.Println("Minimum travel cost:", minCost)
}
