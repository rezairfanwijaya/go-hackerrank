package cutthesticks_test

import (
	cut "hackerrank/algorithm/cut-the-sticks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutTheSticks(t *testing.T) {
	testCases := []struct {
		Name        string
		Arr         []int32
		Expectation []int32
	}{
		{
			Name:        "Test 1",
			Arr:         []int32{5, 4, 4, 2, 2, 8},
			Expectation: []int32{6, 4, 2, 1},
		}, {
			Name:        "Test 2",
			Arr:         []int32{1, 2, 3, 4, 3, 3, 2, 1},
			Expectation: []int32{8, 6, 4, 1},
		}, {
			Name:        "Test 3",
			Arr:         []int32{1, 2, 3},
			Expectation: []int32{3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := cut.CutTheSticks(testCase.Arr)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func TestFindMinElement(t *testing.T) {
	testCases := []struct {
		Name        string
		Arr         []int32
		Expectation int32
	}{
		{Name: "1,2,3", Arr: []int32{1, 2, 3}, Expectation: 1},
		{Name: "10,0,4,1,5", Arr: []int32{10, 0, 4, 1, 5}, Expectation: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := cut.FindMinElement(testCase.Arr)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}

}

func TestRemoveElementInMiddleArray(t *testing.T) {
	tmp := cut.RemoveElementInMiddleArray([]int32{1, 2, 3, 0, 6, 3, 5}, 3)
	t.Log(tmp)
}

func TestMinusArrayWithMinElement(t *testing.T) {
	result := cut.MinusArrayWithMin([]int32{5, 4, 4, 2, 2, 8})
	t.Log(result)
}
