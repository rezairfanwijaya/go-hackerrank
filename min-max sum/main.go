package main

import (
	"fmt"
)

func minMaxSum(arr []int32) {
	var min int32 = arr[0]
	var max int32 = arr[0]
	var total int64 = 0

	for _, value := range arr {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}

		total += int64(value)
	}

	fmt.Println(total-int64(max), total-int64(min))
}

func main() {
	// ada 5 array positif
	// arr[]int={1,2,3,4,8}
	// cari minimal penjumlahan = 1+2+3+4=10
	// cari maksimal penjumlahan = 2+3+4+8=17

	arr := []int32{1, 2, 3, 4, 8}
	minMaxSum(arr)

}
