package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	output := make([]int, 0, 1e4)
	output = append(output, root.Val)
	for i := 0; i < len(root.Children); i++ {
		output = append(output, preorder(root.Children[i])...)
	}
	return output
}
