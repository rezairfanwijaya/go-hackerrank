package main

import "log"

func birthdayCakeCandles(candles []int32) int32 {
	candlesTallest := make(map[int]int)
	max := 0
	length := len(candles)

	for i := 0; i < length-1; i++ {
		if candles[i] >= candles[i+1] {
			if candles[i] >= int32(max) {
				max = int(candles[i])
				candlesTallest[max] += 1
			}
		}
	}

	if candles[length-1] == int32(max) {
		candlesTallest[max] += 1
	}

	return int32(candlesTallest[max])
}

func main() {
	candles := []int32{4, 4, 1, 4, 4}
	log.Println(birthdayCakeCandles(candles))
}
