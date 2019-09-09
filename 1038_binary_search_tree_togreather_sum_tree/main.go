package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	return bst2Gst(&sum, root)
}

func bst2Gst(sum *int, root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	bst2Gst(sum, root.Right)
	*sum = root.Val + *sum
	root.Val = *sum
	bst2Gst(sum, root.Left)

	return root
}
