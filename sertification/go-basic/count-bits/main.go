package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	res := numToBit(126)
	log.Println(res)
}

func numToBit(num int32) int32 {
	res := strconv.FormatInt(int64(num), 2)
	resSplitted := strings.Split(res, "")
	totalBits := 0
	for _, item := range resSplitted {
		if item == "1" {
			totalBits++
		}
	}

	return int32(totalBits)
}
