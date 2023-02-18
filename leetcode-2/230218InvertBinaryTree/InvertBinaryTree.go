package main //https://leetcode.com/problems/invert-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
func main() {

}