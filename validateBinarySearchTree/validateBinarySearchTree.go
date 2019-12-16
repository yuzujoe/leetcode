package main

// TreeNode definition for a binary tree node
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isValidBST(root *TreeNode) bool {
	p := root
	stack := []*TreeNode{}
	sortValues := []int{}

	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(sortValues) != 0 && p.Val <= sortValues[len(sortValues)-1] {
				return false
			}
			sortValues = append(sortValues, p.Val)
			p = p.Right
		}
	}
	return true
}
