package main

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0, 40)
	for i := left; i <= right; i++ {
		j := i
		for ; j > 0; j /= 10 {
			if (j%10 == 0) || (i%(j%10)) != 0 {
				break
			}
		}

		if j == 0 {
			result = append(result, i)
		}
	}

	return result
}
