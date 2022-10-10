package main

import "log"

func getMoneySpent(keyboards, drives []int32, b int32) int32 {
	var tmp []int

	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			sum := keyboards[i] + drives[j]
			log.Println(sum)
			if sum < b {
				tmp = append(tmp, int(sum))
			}
		}
	}

	sum := 0
	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i] > tmp[i+1] {
			sum = tmp[i]
		} else if tmp[i+1] > tmp[i] {
			sum = tmp[i+1]
		}
	}

	if sum == 0 {
		return -1
	}

	log.Println(tmp)
	return int32(sum)
}

func main() {
	keyboards := []int32{3, 1}
	drives := []int32{5, 2, 8}
	b := 10
	log.Println(getMoneySpent(keyboards, drives, int32(b)))
}
