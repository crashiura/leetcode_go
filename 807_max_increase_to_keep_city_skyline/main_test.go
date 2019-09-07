package main

import (
	"testing"
)

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	testData := [][]int{
		[]int{3, 0, 8, 4},
		[]int{2, 4, 5, 7},
		[]int{9, 2, 6, 3},
		[]int{0, 3, 1, 0},
	}

	result := maxIncreaseKeepingSkyline(testData)
	if result != 35 {
		t.Errorf("Error expected %d, got %d", result, 35)
	}
}
