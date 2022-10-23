package main

import (
	"log"
	"math"
)

func saveThePrisoner(n, m, s int32) int32 {
	remain := math.Abs(float64(n - s + 1))
	tmp := math.Abs(float64(m - int32(remain)))
	res := math.Abs(float64(int32(tmp) - n))
	return int32(res)
}

func main() {
	log.Println(saveThePrisoner(5, 2, 1))
}
