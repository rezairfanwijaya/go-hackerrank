package main

import (
	"log"
)

func compareTriplets(a, b []int32) []int32 {
	tmp := make(map[string]int)

	for index := range a {
		if a[index] > b[index] {
			tmp["a"] += 1
		} else if a[index] < b[index] {
			tmp["b"] += 1
		}
	}

	var res []int32
	res = append(res, int32(tmp["a"]))
	res = append(res, int32(tmp["b"]))

	return res
}

func main() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}

	log.Println(compareTriplets(a, b))
}
