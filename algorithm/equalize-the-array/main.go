package main

func equalizeArray(arr []int32) int32 {
	// kumpulkan semua angka dan hitung ada berapa banyaknya
	dictionary := make(map[int32]int32)
	for _, v := range arr {
		// tambahkan key baru ketika kelum ada
		if _, ok := dictionary[v]; !ok {
			dictionary[v] = 1
		} else {
			// increment value ketiak sudah ada
			dictionary[v] += 1
		}
	}

	// hitung value yang paling banyak
	var max int32 = 0
	for _, v := range dictionary {
		if v > max {
			max = v
		}
	}

	// kurangi panjang array dengan nilai max
	return int32(len(arr)) - max
}

func main() {
	arr := []int32{3, 3, 2, 1, 3}
	println(equalizeArray(arr))
}
