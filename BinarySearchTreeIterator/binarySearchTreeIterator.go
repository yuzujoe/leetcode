package main

// TreeNode  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BSTIterator struct
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor defini constructor
func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}
	return BSTIterator{
		stack: stack,
	}
}

// Next return the next smallest number
func (this *BSTIterator) Next() int {
	cur := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	ret := cur.Val
	cur = cur.Right
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}

	return ret
}

// HasNext return whether we have a next smallest number
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
