package main

import (
	"log"
	"math"
)

func lowestTriangle(trianglebase, area int32) {
	tb := float64(trianglebase)
	a := float64(area)

	result := math.Round(2 * a / tb)

	if result == 0 {
		log.Println(1)
	} else {
		log.Println(result)
	}
}

func main() {
	lowestTriangle(1000000, 1)
}
