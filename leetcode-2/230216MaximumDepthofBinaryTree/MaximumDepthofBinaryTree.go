package main //https://leetcode.com/problems/maximum-depth-of-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxOf(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxOf(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {

}
