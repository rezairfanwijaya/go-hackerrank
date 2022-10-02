package main

import "log"

func veryBigSum(arr []int64) int64 {
	var sum int64

	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	log.Println(veryBigSum(arr))
}
