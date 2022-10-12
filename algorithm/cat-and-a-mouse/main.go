package main

import (
	"log"
	"math"
)

func catAndMouse(x, y, z int32) string {
	catA := 0
	catB := 0

	catA = int(math.Abs(float64(z - x)))
	catB = int(math.Abs(float64(z - y)))

	if catA < catB {
		return "Cat A"
	} else if catA > catB {
		return "Cat B"
	}

	return "Mouse C"
}

func main() {
	x := 1
	y := 3
	z := 2

	log.Println(catAndMouse(int32(x), int32(y), int32(z)))
}
