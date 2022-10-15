package main

import "log"

func handShake(n int32) int32 {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	var array []int
	for i := 1; i <= int(n); i++ {
		array = append(array, i)
	}

	shake := 0
	index := 1
	for i := index; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[j] == array[len(array)-1] {
				index++
			}
			shake++
		}
	}

	return int32(shake)
}

func main() {
	log.Println(handShake(5))
}
