package main

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Left != nil {
		if isSameOrAncestorNode(root.Left, p) && isSameOrAncestorNode(root.Left, q) {
			return lowestCommonAncestor(root.Left, p, q)
		}
	}

	if root.Right != nil {
		if isSameOrAncestorNode(root.Right, p) && isSameOrAncestorNode(root.Right, q) {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}

	return root
}

func isSameOrAncestorNode(root *TreeNode, node *TreeNode) bool {
	if root == node {
		return true
	}

	if root.Left != nil && isSameOrAncestorNode(root.Left, node) {
		return true
	}

	if root.Right != nil && isSameOrAncestorNode(root.Right, node) {
		return true
	}

	return false
}
