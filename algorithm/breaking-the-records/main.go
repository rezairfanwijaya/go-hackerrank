package main

import "log"

func breakingTheRecord(scores []int32) []int32 {
	highestScores := scores[0]
	lowestScores := scores[0]

	breakHigestScores := 0
	breakLowestScores := 0

	for _, score := range scores {
		if score > highestScores {
			breakHigestScores++
			highestScores = score
		}

		if score < lowestScores {
			breakLowestScores++
			lowestScores = score
		}
	}

	return []int32{int32(breakHigestScores), int32(breakLowestScores)}
}

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	log.Println(breakingTheRecord(scores))
}
