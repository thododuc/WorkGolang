package main //https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	for root.Val > q.Val || root.Val < p.Val {
		if root.Val > q.Val {
			root = root.Left
			continue
		}
		if root.Val < p.Val {
			root = root.Right
		}
	}
	return root
}
func main() {

}
