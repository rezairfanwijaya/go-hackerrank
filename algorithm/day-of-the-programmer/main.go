package main

import (
	"fmt"
	"log"
)

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return fmt.Sprintf("26.09.%d", year)
	}

	if year >= 1700 && year <= 1917 && year%4 == 0 {
		return fmt.Sprintf("12.09.%d", year)
	}

	if year >= 1919 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		return fmt.Sprintf("12.09.%d", year)
	}

	return fmt.Sprintf("13.09.%d", year)
}

func main() {
	year := 2400
	log.Println(dayOfProgrammer(int32(year)))
}
