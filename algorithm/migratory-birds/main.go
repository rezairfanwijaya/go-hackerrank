package main

import "log"

func migratoryBirds(arr []int32) int32 {
	birdIds := make(map[int]int)

	for _, value := range arr {
		_, isExist := birdIds[int(value)]
		if isExist {
			birdIds[int(value)] += 1
		} else {
			birdIds[int(value)] = 1
		}
	}

	max := 0
	birdId := 0
	for index, value := range birdIds {
		if value > max {
			max = value
			birdId = index
		}
	}

	return int32(birdId)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4, 4, 4, 4, 4, 4}
	log.Println(migratoryBirds(arr))
}
