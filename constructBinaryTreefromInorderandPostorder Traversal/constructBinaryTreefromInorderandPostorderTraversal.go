package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	nodeIndex := findIndex(rootVal, inorder)
	Node := &TreeNode{rootVal, nil, nil}

	Node.Left = buildTree(inorder[:nodeIndex], postorder[:nodeIndex])
	Node.Right = buildTree(inorder[nodeIndex+1:], postorder[nodeIndex:len(postorder)-1])

	return Node
}

func findIndex(expected int, arr []int) int {
	for i, v := range arr {
		if v == expected {
			return i
		}
	}
	return -1
}
