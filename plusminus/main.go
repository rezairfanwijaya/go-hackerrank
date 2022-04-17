package main

import "fmt"

func PlusMinus(arr []int32) {
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
	arr := []int32{2, 2, 3, 3, 0, -4}
	PlusMinus(arr)
}
