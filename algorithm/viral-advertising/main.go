package main

import "log"

func viralAdvertising(n int32) int32 {
	progress := make(map[int][]int32)

	for i := 1; i <= int(n); i++ {
		if i == 1 {
			progress[i] = []int32{5, 2, 2}
		} else {
			nextDay(progress, i)
		}
	}

	return progress[int(n)][2]
}

func nextDay(progress map[int][]int32, i int) {
	index := progress[i-1]
	shared := index[0] / 2 * 3
	liked := int(shared / 2)
	cumulative := index[2] + int32(liked)
	progress[i] = []int32{int32(shared), int32(liked), cumulative}
}

func main() {
	log.Println(viralAdvertising(3))
}
