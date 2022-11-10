package repeatedstring_test

import (
	repeatedstring "hackerrank/algorithm/repeated-string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonDivisbleSubset(t *testing.T) {
	testCases := []struct {
		Name        string
		Sentence    string
		Counter     int64
		Expectation int64
	}{
		{
			Name:        "Test 1",
			Sentence:    "abcac",
			Counter:     10,
			Expectation: 4,
		}, {
			Name:        "Test 2",
			Sentence:    "aba",
			Counter:     10,
			Expectation: 7,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := repeatedstring.RepetedString(testCase.Sentence, testCase.Counter)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}

}
