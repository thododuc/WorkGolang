// https://leetcode.com/problems/validate-binary-search-tree
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MaxInt, math.MinInt)
}

func validate(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	} else {
		if root.Val <= min || root.Val >= max {
			return false
		} else {
			return validate(root.Left, root.Val, min) && validate(root.Right, max, root.Val)
		}
	}
}
