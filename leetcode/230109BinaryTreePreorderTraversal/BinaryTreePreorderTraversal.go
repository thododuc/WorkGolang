package main //https://leetcode.com/problems/binary-search

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}
