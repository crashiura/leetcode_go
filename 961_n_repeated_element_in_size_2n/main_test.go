package main

import (
	"testing"
)

func Test(t *testing.T) {
	data := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{1, 2, 3, 3},
			out: 3,
		},
		{
			in:  []int{2, 1, 2, 5, 3, 2},
			out: 2,
		},
		{
			in:  []int{5, 1, 5, 2, 5, 3, 5, 4},
			out: 5,
		},
	}

	for _, v := range data {
		res := repeatedNTimes(v.in)
		if res != v.out {
			t.Errorf("Error expected %d, got %d", res, v.out)
		}
	}

}
