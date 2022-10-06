package main

import "log"

func birthDays(s []int32, d, m int32) int32 {

	length := len(s)
	var totalDivide int32 = 0
	var sum int32 = 0
	var counter int32 = 0

	if s[0] == d || length == 1 {
		return 1
	}

	for i := 0; i < length; i++ {

		limit := m + int32(i)
		if int(limit) <= len(s) {
			slice := s[i : m+int32(i)]

			for _, v := range slice {
				sum += v
				if sum == d {
					totalDivide++
					sum = 0
					counter = 0
				}

				if counter == m {
					sum = 0
					counter = 0
				}
				counter++
			}
		}
	}

	return totalDivide
}

func main() {

	s := []int32{1, 2, 1, 3, 2}
	// birthDays(s, 3, 2)
	log.Println(birthDays(s, 3, 2))

}
