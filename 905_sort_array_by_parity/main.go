package main

import (
	"sort"
)

func sortArrayByParity(A []int) []int {
	odd := make([]int, 0, 20)
	even := make([]int, 0, 20)

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			even = append(even, A[i])
			continue
		}
		odd = append(odd, A[i])
	}
	sort.Ints(odd)
	sort.Ints(even)
	result := make([]int, 0, len(even)+len(odd))
	result = append(even, odd...)
	return result
}
