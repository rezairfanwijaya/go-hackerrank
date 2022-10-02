package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	var plus []int32
	var minus []int32
	var netral []int32

	for _, v := range arr {
		if v > 0 {
			plus = append(plus, v)
		} else if v < 0 {
			minus = append(minus, v)
		} else {
			netral = append(netral, v)
		}
	}

	fmt.Println(float64(len(plus)) / float64(len(arr)))
	fmt.Println(float64(len(minus)) / float64(len(arr)))
	fmt.Println(float64(len(netral)) / float64(len(arr)))

}

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}
