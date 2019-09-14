package main

import "testing"

func TestJudgeCircle(t *testing.T) {
	tcs := []struct {
		moves  string
		result bool
	}{
		{
			moves:  "UD",
			result: true,
		},
		{
			moves:  "LL",
			result: false,
		},
	}

	for _, tc := range tcs {
		res := judgeCircle(tc.moves)
		if res != tc.result {
			t.Errorf("Error expected %t, got %t", res, tc.result)
		}
	}
}
