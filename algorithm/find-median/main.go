package main

import (
	"fmt"
)

func findMedian(arr []int32) int32 {
	// sort arrays
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	indexMedian := len(arr) / 2
	return arr[indexMedian]
}

func main() {
	fmt.Println(findMedian([]int32{0, 1, 2, 4, 6, 5, 3}))
}
