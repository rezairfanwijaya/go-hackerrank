package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChocolateFeast(t *testing.T) {
	testCases := []struct {
		Name          string
		Amount        int32
		ChocoPrice    int32
		WrapperNumber int32
		ExpectedRes   int32
	}{
		{
			Name:          "success 1",
			Amount:        15,
			ChocoPrice:    3,
			WrapperNumber: 2,
			ExpectedRes:   9,
		},
		{
			Name:          "success 1",
			Amount:        12,
			ChocoPrice:    4,
			WrapperNumber: 4,
			ExpectedRes:   3,
		},
		{
			Name:          "success 1",
			Amount:        6,
			ChocoPrice:    2,
			WrapperNumber: 2,
			ExpectedRes:   5,
		},
	}

	for _, tc := range testCases {
		res := ChocolateFeast(tc.Amount, tc.ChocoPrice, tc.WrapperNumber)
		require.Equal(t, tc.ExpectedRes, res)
	}
}
