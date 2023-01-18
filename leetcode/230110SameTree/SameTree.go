package main //https://leetcode.com/problems/same-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	v := false
	switch {
	case p == nil && q == nil:
		return true
	case p == nil && q != nil:
		return false
	case p != nil && q == nil:
		return false
	default:
		if p.Val == q.Val {
			v = true
		} else {
			return false
		}
		l := isSameTree(p.Left, q.Left)
		r := isSameTree(p.Right, q.Right)
		return v && l && r
	}
}
