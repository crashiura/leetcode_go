package main

func repeatedNTimes(A []int) int {
	store := make(map[int]struct{}, 10000)
	for _, v := range A {
		if _, found := store[v]; found {
			return v
		}

		store[v] = struct{}{}
	}

	return -1
}
