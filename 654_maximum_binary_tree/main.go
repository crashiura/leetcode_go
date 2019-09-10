package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	max := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			maxIndex = i
		}
	}

	root := TreeNode{Val: nums[maxIndex]}
	if maxIndex >= 1 {
		root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	}
	if maxIndex+1 < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	}
	return &root
}
