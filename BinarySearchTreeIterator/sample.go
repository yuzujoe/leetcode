package main

// TreeNode  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nodes := []*TreeNode{}
	node := root
	for node != nil {
		nodes = append(nodes, node)
		node = node.Left
	}
	return BSTIterator{nodes: nodes}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	cur := node.Right
	for cur != nil {
		this.nodes = append(this.nodes, cur)
		cur = cur.Left
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}
