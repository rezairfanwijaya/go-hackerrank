package main

import "log"

func divisibleSumPairs(n, k int32, ar []int32) int32 {

	sumPairs := 0

	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[i]+ar[j+1] == k {
				sumPairs++
			}
		}
	}

	return int32(sumPairs)
}

func main() {
	ar := []int32{1, 3, 2, 6, 1, 2}
	log.Println(divisibleSumPairs(6, 3, ar))

}
