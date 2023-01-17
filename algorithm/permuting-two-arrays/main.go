package main

import (
	"fmt"
)

func twoArrays(k int32, A, B []int32) string {
	// set terlenbih dahulu variable penentu
	// untuk bisa permutasi atau tidak
	// dengan devault value adalah false
	isPermutation := false

	// jadikan slice A  menjadi map sebagai
	// patokan
	var base = make(map[int32]int32)
	// satu angka hanya bisa berpasangan
	// dengan satu angka
	var isUsed = make(map[int32]bool)
	for key, v := range A {
		base[int32(key)] = v
		isUsed[int32(key)] = false
	}

	// loop B dan cocokan apakah ketika ditambah dengan
	// base akan menjadi K ?
	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B); j++ {
			thisBase := base[int32(j)]
			if B[i]+thisBase == k {
				// cek apakah base sudah dipakai
				if !isUsed[thisBase] {
					isPermutation = true
					isUsed[thisBase] = true
				} else {
					isPermutation = false
				}
			}
		}
	}

	if isPermutation {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(twoArrays(10, []int32{2, 1, 3}, []int32{7, 8, 9}))
}
