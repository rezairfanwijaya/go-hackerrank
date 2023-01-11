package main

import (
	"fmt"
	"strconv"
)

func flippingBits(n int64) int64 {
	// variable bit dalam string
	bitOfNInString := fmt.Sprintf("%032b", n)

	// ubah 0=>1
	// ubah 1=>0

	// variable bit setelah di replace
	var bitOfNAfterReplace string

	// proses replace
	for i := 0; i < len(bitOfNInString); i++ {
		tmp := string(bitOfNInString[i])
		if tmp == "0" {
			bitOfNAfterReplace += "1"
		} else {
			bitOfNAfterReplace += "0"
		}
	}

	// ubah binary yang sudah direplace
	// kedalam int64
	result, _ := strconv.ParseInt(bitOfNAfterReplace, 2, 64)

	return result
}

func main() {
	fmt.Println(flippingBits(2147483647))
}
