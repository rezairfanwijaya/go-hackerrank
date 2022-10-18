package main

import "log"

func utopianTree(n int32) int32 {
	period := make(map[int]int32)
	height := 1

	// set period
	for i := 0; i <= 60; i++ {
		if i == 0 {
			period[i] = int32(height)
			height++
		} else if i%2 != 0 && i != 0 {
			before := period[i-1]
			period[i] = before * 2
			height = int(period[i])
		} else {
			period[i] = int32(height) + 1
		}
	}

	res := period[int(n)]
	return res
}

func main() {
	log.Println(utopianTree(0))
}
