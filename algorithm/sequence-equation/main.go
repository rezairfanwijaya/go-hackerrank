package main

import "log"

func permutaionEquation(p []int32) []int32 {
	tmp := make(map[int32]int32)
	var result []int32
	key := 1

	// pairing element p dengan konstanta
	for i := 0; i < len(p); i++ {
		tmp[p[i]] = int32(key)
		key++
	}

	// ambil y
	for i := 1; i <= len(p); i++ {
		x := tmp[int32(i)]
		y := tmp[x]
		result = append(result, y)
	}

	return result
}

func main() {
	p := []int32{4, 3, 5, 1, 2}
	log.Println(permutaionEquation(p))
}
