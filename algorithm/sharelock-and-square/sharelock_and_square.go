package sharelockandsquare

import (
	"math"
)

func Squares(a, b int32) int32 {
	var result int32

	for i := a; i <= b; i++ {
		tmp := math.Sqrt(float64(i))
		x := int32(tmp)
		if x*x == i {
			result++
		}
	}

	return result
}
