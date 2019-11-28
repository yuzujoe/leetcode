package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	levelOrder := breadthFirst(getChildren, root)
	for _, level := range levelOrder {
		var levelValues []int
		for _, item := range level {
			levelValues = append(levelValues, item.Val)
		}
		res = append(res, levelValues)
	}
	return res
}

func breadthFirst(f func(n *TreeNode) []*TreeNode, root *TreeNode) [][]*TreeNode {
	var res [][]*TreeNode
	level := []*TreeNode{root}
	for len(level) > 0 {
		res = append(res, level)
		var nextLevel []*TreeNode
		for _, item := range level {
			nextLevel = append(nextLevel, f(item)...)
		}
		level = nextLevel
	}
	return res
}

func getChildren(n *TreeNode) []*TreeNode {
	var res []*TreeNode
	if n.Left != nil {
		res = append(res, n.Left)
	}
	if n.Right != nil {
		res = append(res, n.Right)
	}
	return res
}

func main() {

}
