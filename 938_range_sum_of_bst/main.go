package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L, R int) int {
	if root == nil {
		return 0
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
