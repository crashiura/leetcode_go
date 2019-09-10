package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insertInTree(root, val)
	return root
}

func insertInTree(root *TreeNode, val int) {
	if val > root.Val {
		rNode := root.Right
		if rNode != nil {
			insertIntoBST(rNode, val)
		} else {
			root.Right = &TreeNode{Val: val}
		}
		return
	}

	if val < root.Val {
		lNode := root.Left
		if lNode != nil {
			insertIntoBST(lNode, val)
		} else {
			root.Left = &TreeNode{Val: val}
		}
	}
}
