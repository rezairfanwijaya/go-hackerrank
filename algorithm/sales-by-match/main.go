package main

import "log"

func sockMerchant(n int32, ar []int32) int32 {
	sock := make(map[int]int)

	for _, value := range ar {
		_, isExist := sock[int(value)]
		if isExist {
			sock[int(value)] += 1
		} else {
			sock[int(value)] = 1
		}
	}

	match := 0
	for _, v := range sock {
		match += v / 2

	}

	return int32(match)
}

func main() {
	n := 7
	ar := []int32{1, 2, 1, 2, 1, 3, 2}
	log.Println(sockMerchant(int32(n), ar))
}
