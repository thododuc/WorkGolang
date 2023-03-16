package main //https://leetcode.com/problems/path-sum/description/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	} else {
		return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
	}
}