package main

import (
	"fmt"
)

func matchingStrings(strings, queries []string) []int32 {
	// bikin kamus dari strings dan queries
	mapStrings := make(map[string]int)
	var mapQueries []int32

	for _, v := range strings {
		// cek apakah ada key v kalau tidak ada maka add
		if _, ok := mapStrings[v]; !ok {
			mapStrings[v] = 1
		} else {
			mapStrings[v] += 1
		}
	}

	for _, v := range queries {
		// cek apakah query ada pada mapString, kalau ada ambil value nya
		if value, ok := mapStrings[v]; ok {
			mapQueries = append(mapQueries, int32(value))
		} else {
			mapQueries = append(mapQueries, 0)
		}
	}

	return mapQueries
}

func main() {
	fmt.Println(matchingStrings([]string{"ab", "ab", "jkl", "ab", "abc"}, []string{"ab", "abc", "bc", "jkl"}))
}
