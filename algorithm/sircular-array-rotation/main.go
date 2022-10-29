package main

import "log"

// one problem for this logic is time complexity O(n^2)
func sircularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var result []int32

	for i := 0; i < int(k); i++ {
		var x []int32
		// taruh last index di depan
		x = append(x, a[len(a)-1])

		// geser dan gabung element selain last index
		for j := 0; j < len(a)-1; j++ {
			x = append(x, a[j])
		}

		// update array
		a = x
	}

	// akses index di lastest array a berdasarkan query
	for _, query := range queries {
		result = append(result, a[query])
	}

	return result
}

func main() {
	a := []int32{3, 4, 5, 6}
	k := 1
	queries := []int32{1, 2}

	log.Println(sircularArrayRotation(a, int32(k), queries))
}
