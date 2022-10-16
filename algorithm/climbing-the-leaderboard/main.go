package main

import (
	"log"
)

func climbingLeaderBoard(ranked, player []int32) []int32 {
	rankList := make(map[int]int)
	rank := 1

	// get ranklist
	for i := 0; i < len(ranked); i++ {
		_, isExist := rankList[int(ranked[i])]
		if isExist {
			rank--
			rankList[int(ranked[i])] = rank
			rank++
		} else {
			rankList[int(ranked[i])] = rank
			rank++
		}
	}

	// highest lowest for rank
	highest := 0
	lowest := 0
	// distance := 0
	log.Println(rankList)

	for _, rank := range ranked {
		if rank > int32(highest) {
			highest = int(rank)
		} else {
			lowest = int(rank)
		}
	}

	var res []int32

	for _, player := range player {
		if player < int32(lowest) {
			res = append(res, int32(rankList[lowest]+1))
		} else if player > int32(highest) {
			res = append(res, int32(rankList[highest]))
		} else {
			_, isExist := rankList[int(player)]
			if isExist {
				res = append(res, int32(rankList[int(player)]))
			} else {
				// complete this
			}
		}
	}

	return res
}

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}
	log.Println(climbingLeaderBoard(ranked, player))
}
