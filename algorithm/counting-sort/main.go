package main

import (
	"fmt"
)

func countingSort(arr []int32) []int32 {

	// membuat array 0 sejumlah panjang arr
	countSort := make([]int32, 100)

	// loop arr dan ubah 0 value sesuai dengan value arr
	for _, v := range arr {
		countSort[v]++
	}

	return countSort
}

func main() {
	fmt.Println(countingSort([]int32{1, 1, 3, 2, 1}))
}
