package main

import (
	"math"
)

func minimumDistances(a []int32) int32 {
	pairs := make(map[int32][]int32)

	for i := 0; i < len(a); i++ {
		pairs[a[i]] = append(pairs[a[i]], int32(i))
	}

	var minimumDistance int32

	for _, v := range pairs {
		if len(v) >= 2 {
			distance := getDistance(v)
			if minimumDistance == 0 {
				minimumDistance = distance
			} else {
				if minimumDistance > distance {
					minimumDistance = distance
				}
			}

		}
	}

	if minimumDistance == 0 {
		return -1
	}

	return int32(minimumDistance)
}

func getDistance(pairs []int32) int32 {
	var distance int32

	for index, v := range pairs {
		if index == 0 {
			distance = pairs[0]
		} else {
			distance -= v
		}
	}

	return int32(math.Abs(float64(distance)))
}

func main() {
	minimumDistances([]int32{1, 2, 3, 1, 8, 4, 2})
	// getDistance([]int32{1, 5})
}
