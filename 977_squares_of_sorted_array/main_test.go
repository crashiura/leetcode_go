package main

import (
	"reflect"
	"testing"
)

func TestSquaresOfSortedArray(t *testing.T) {
	tcs := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{-4, -1, 0, 3, 10},
			out: []int{0, 1, 9, 16, 100},
		},
		{
			in:  []int{-7, -3, 2, 3, 11},
			out: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range tcs {
		res := sortedSquares(tc.in)
		if reflect.DeepEqual(res, tc.out) == false {
			t.Errorf("Expected %v, got %v", res, tc.out)
		}
	}
}
