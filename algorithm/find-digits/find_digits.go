package finddigits

import (
	"strconv"
)

func FindDigits(n int32) int32 {
	// initial divider
	var divider int32
	// change n to string for looping
	numberString := strconv.Itoa(int(n))

	// find digits
	for _, v := range numberString {
		tmp := string(v)
		number, _ := strconv.Atoi(tmp)
		// find divider
		if number != 0 {
			if n%int32(number) == 0 {
				divider++
			}
		}
	}

	return divider
}
