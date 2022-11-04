package sharelockandsquare_test

import (
	sharelockandsquare "hackerrank/algorithm/sharelock-and-square"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquares(t *testing.T) {
	testCases := []struct {
		Name     string
		A        int32
		B        int32
		Expected int32
	}{
		{
			Name:     "24 49",
			A:        24,
			B:        49,
			Expected: 3,
		}, {
			Name:     "3 9",
			A:        3,
			B:        9,
			Expected: 2,
		}, {
			Name:     "17 24",
			A:        17,
			B:        24,
			Expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := sharelockandsquare.Squares(testCase.A, testCase.B)
			assert.Equal(t, testCase.Expected, actual)
		})
	}
}
