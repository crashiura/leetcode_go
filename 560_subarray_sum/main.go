package main

import "fmt"

func main() {
	fmt.Println(subArraySum2([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0))
}

func subArraySum(nums []int, k int) int {
	cnt := 0

	for start := 0; start < len(nums); start++ {
		sum := 0

		for j := start; j < len(nums); j++ {

			sum += nums[j]

			if sum == k {
				cnt++
			}
		}
	}

	return cnt
}

func subArraySum2(nums []int, k int) int {
	cnt, sum := 0, 0
	Map := make(map[int]int)
	Map[0] = 1

	for _, n := range nums {
		if n == 0 && k == 0 {
			cnt++
			continue
		}

		sum += n
		cnt += Map[sum-k]
		Map[sum]++
	}

	return cnt
}

func subarraySum(nums []int, k int) int {
	sumMap := map[int]int{0: 1}
	sum := 0
	res := 0
	for _, n := range sumMap {
		if n == 0 && k == 0 {
			res++
			continue
		}

		sum += n
		res += sumMap[sum-k]
		sumMap[sum]++
	}

	return res
}
