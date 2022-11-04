package finddigits_test

import (
	finddigits "hackerrank/algorithm/find-digits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDigits(t *testing.T) {
	testCases := []struct {
		Name     string
		Number   int32
		Expected int32
	}{
		{
			Name:     "111",
			Number:   111,
			Expected: 3,
		},
		{
			Name:     "124",
			Number:   124,
			Expected: 3,
		},
		{
			Name:     "10",
			Number:   10,
			Expected: 1,
		},
		{
			Name:     "12",
			Number:   12,
			Expected: 2,
		},
		{
			Name:     "1012",
			Number:   1012,
			Expected: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := finddigits.FindDigits(testCase.Number)
			assert.Equal(t, testCase.Expected, actual)
		})
	}
}
