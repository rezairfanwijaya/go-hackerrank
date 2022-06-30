package main

import "fmt"

func countBit(num uint32) int32 {
	var count uint32 = 0
	for num != 0 {
		count += num & 1
		num >>= 1
	}

	return int32(count)
}

func main() {
	bit := countBit(7)
	fmt.Println(bit)
}
