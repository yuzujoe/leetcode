package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return rec(root, nil, nil)
}

func rec(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}

	left := rec(node.Left, min, &node.Val)
	right := rec(node.Right, &node.Val, max)

	if !left {
		return false
	}
	return right
}
