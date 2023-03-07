package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDistances(t *testing.T) {
	testCases := []struct {
		Name        string
		Numbers     []int32
		Expectation int32
	}{
		{
			Name:        "success with pairs",
			Numbers:     []int32{1, 2, 3, 1, 8, 4, 2},
			Expectation: 3,
		}, {
			Name:        "success with no pairs",
			Numbers:     []int32{1, 2, 3, 88, 8, 4, 90},
			Expectation: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			minimumDistance := minimumDistances(testCase.Numbers)
			assert.Equal(t, testCase.Expectation, minimumDistance)
		})
	}
}

func TestGetDistance(t *testing.T) {
	testCases := struct {
		Name        string
		Numbers     []int32
		Expectation int32
	}{
		Name:        "success",
		Numbers:     []int32{1, 2},
		Expectation: 1,
	}

	t.Run(testCases.Name, func(t *testing.T) {
		distance := getDistance(testCases.Numbers)
		assert.Equal(t, testCases.Expectation, distance)
	})
}
