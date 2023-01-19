package main //https://leetcode.com/problems/symmetric-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}
func isMirror(rootLeft *TreeNode, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return true
	} else if rootLeft == nil || rootRight == nil {
		return false
	} else {
		if rootLeft.Val == rootRight.Val {
			return isMirror(rootLeft.Left, rootRight.Right) && isMirror(rootLeft.Right, rootRight.Left)
		} else {
			return false
		}
	}
}
