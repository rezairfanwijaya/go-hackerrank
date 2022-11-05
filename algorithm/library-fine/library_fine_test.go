package libraryfine_test

import (
	libraryfine "hackerrank/algorithm/library-fine"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibraryFine(t *testing.T) {
	testCases := []struct {
		Name     string
		D1       int32
		M1       int32
		Y1       int32
		D2       int32
		M2       int32
		Y2       int32
		Expected int32
	}{
		{
			Name:     "No late",
			D1:       1,
			M1:       1,
			Y1:       2000,
			D2:       1,
			M2:       1,
			Y2:       2000,
			Expected: 0,
		}, {
			Name:     "Late year",
			D1:       2,
			M1:       8,
			Y1:       2003,
			D2:       3,
			M2:       5,
			Y2:       2000,
			Expected: 10000,
		}, {
			Name:     "Late month",
			D1:       12,
			M1:       4,
			Y1:       2012,
			D2:       11,
			M2:       2,
			Y2:       2012,
			Expected: 1000,
		}, {
			Name:     "Late day",
			D1:       22,
			M1:       7,
			Y1:       2021,
			D2:       12,
			M2:       7,
			Y2:       2021,
			Expected: 150,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := libraryfine.LibraryFine(
				testCase.D1,
				testCase.M1,
				testCase.Y1,
				testCase.D2,
				testCase.M2,
				testCase.Y2,
			)
			assert.Equal(t, testCase.Expected, actual)
		})
	}
}
