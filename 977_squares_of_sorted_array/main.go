package main

import (
	"sort"
)

func sortedSquares1(A []int) []int {
	for i, v := range A {
		A[i] = v * v
	}

	sort.Ints(A)
	return A
}
