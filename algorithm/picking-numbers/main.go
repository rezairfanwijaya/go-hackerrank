package main

import (
	"log"
	"math"
)

func pickingNumbers(a []int32) int32 {
	var tmp []int
	tmp = append(tmp, int(a[0]))

	for i := 0; i < len(a)-1; i++ {
		sum := math.Abs(float64(a[0] - a[i+1]))
		distance := math.Abs(float64(int(a[i+1]) - tmp[0]))
		if sum <= 1 && sum >= 0 {
			if distance <= 1 && distance >= 0 {
				tmp = append(tmp, int(a[i+1]))
			}
		}
	}

	res := len(tmp)
	return int32(res)
}

func main() {
	a := []int32{
		4, 6, 5, 3, 3, 1,
	}

	log.Println(pickingNumbers(a))
}
