package main

import (
	"fmt"
	"strconv"
)

func kaprekarNumbers(p, q int32) {
	const (
		TWO_DIGIT = 10
	)

	var res []int64

	var i int32
	for i = p; i <= q; i++ {
		// jika 1 maka langsung append
		if i == 1 {
			res = append(res, 1)
		}

		square := i * i
		// ambil bilangan yang berjumlah 2 digit atau lebih (genap)
		if square >= TWO_DIGIT {
			// split dan sum, hasilnya harus sama dengan i
			if splitAndSum(int64(square)) == int64(i) {
				res = append(res, int64(i))
			}
		}

	}

	if len(res) == 0 {
		fmt.Println("INVALID RANGE")
		return
	}

	for _, v := range res {
		fmt.Printf("%v ", v)
	}
}

func splitAndSum(square int64) int64 {
	// konversi ke string
	squaresring := strconv.Itoa(int(square))

	// cek panjang square
	length := len(squaresring) / 2

	// sisi kiri
	leftSide := squaresring[:length]
	// sisi kanan
	rightSide := squaresring[length:]

	// convert to int
	leftSideInt, _ := strconv.Atoi(leftSide)
	rightSideInt, _ := strconv.Atoi(rightSide)

	return int64(leftSideInt + rightSideInt)
}

func main() {
	kaprekarNumbers(1, 99999)
}
