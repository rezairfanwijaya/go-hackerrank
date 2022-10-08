package main

import (
	"fmt"
)

func bonAppetit(bill []int32, k, b int32) {
	var total int32
	for i := 0; i < len(bill); i++ {
		if bill[i] != bill[k] {
			total += bill[i]
		}
	}

	charged := b - total/2
	if charged == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged)
	}

}

func main() {
	bill := []int32{3, 10, 2, 9}
	k := 1
	b := 12
	bonAppetit(bill, int32(k), int32(b))
}
