package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func beautifullDays(i, j, k int32) int32 {
	beautifulDay := 0

	for z := i; z <= j; z++ {
		reverseNumber := reverseNumber(z)
		tmp := z - reverseNumber
		distance := math.Abs(float64(tmp))
		if int32(distance)%k == 0 {
			beautifulDay++
		}
	}

	return int32(beautifulDay)
}

func reverseNumber(num int32) int32 {
	stringNumber := strconv.Itoa(int(num))
	var tmpStringNumber []string

	for i := len(stringNumber) - 1; i >= 0; i-- {
		tmpStringNumber = append(tmpStringNumber, string(stringNumber[i]))
	}

	join := strings.Join(tmpStringNumber, "")
	reverseNumber, _ := strconv.Atoi(join)

	return int32(reverseNumber)
}

func main() {
	log.Println(beautifullDays(20, 23, 6))
}
