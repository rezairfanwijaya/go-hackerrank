package main

import "log"

func hurdleRace(k int32, height []int32) int32 {
	// ambil max height
	var maxHeight int32 = 0
	for _, v := range height {
		if v > maxHeight {
			maxHeight = v
		}
	}

	// bila jump of character gt max height
	if k > maxHeight {
		return 0
	}

	// selisih jump dengan max height untuk dosis
	return maxHeight - k
}

func main() {
	k := 7
	height := []int32{2, 5, 4, 5, 2}
	log.Println(hurdleRace(int32(k), height))
}
