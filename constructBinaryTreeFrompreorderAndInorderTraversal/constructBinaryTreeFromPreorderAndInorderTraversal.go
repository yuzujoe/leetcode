package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0
	for index < len(inorder) {
		if inorder[index] == preorder[0] {
			break
		}
		index++
	}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
