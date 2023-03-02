package main //https://leetcode.com/problems/squares-of-a-sorted-array

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i:=0; i < len(nums) ; i++ {
		nums[i] *= nums[i]
}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func main() {
	nums := []int {-4,-1,0,3,10}
	fmt.Println(sortedSquares(nums))
}