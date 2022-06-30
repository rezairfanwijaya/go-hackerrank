package main

import "fmt"

func median(arr []int32) int32 {

	swapped := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}

		}

		if !swapped {
			return 0
		}
	}

	return arr[len(arr)/2]

}

func main() {
	num := []int32{5, 6, 7, 3, 4, 2, 1}
	// 1, 2, 3, 4, 5, 6,
	// nums := []int32{1, 2, 3, 4, 5}
	// res := median(num)
	// fmt.Println(res)
	// median(nums)
	fmt.Println(median(num))
}
