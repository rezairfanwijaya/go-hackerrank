package main

import "log"

func arraySum(arr []int32) int32{
	sum := 0
	for _,value := range arr {
		sum += int(value)
	}

	return int32(sum)
} 


func main() {
	arr := []int32{1, 2, 3, 4, 10, 11}
	log.Println(arraySum(arr))
}