package main

import "fmt"

func lonelyInteger(arr []int32) int32 {

	res := make(map[int32]int32)
	for _, item := range arr {
		res[item] = res[item] + 1
	}

	var final int32
	for k, v := range res {
		if v == 1 {
			final = k
		}
	}

	return final

}

func main() {
	arr := []int32{1, 2, 3, 2, 1, 1}
	// fmt.Println(lonelyInteger(arr))
	fmt.Println(lonelyInteger(arr))

}
