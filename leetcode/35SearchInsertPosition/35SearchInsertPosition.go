//https://leetcode.com/problems/search-insert-position/description/

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
	nums1 := []int{}
	fmt.Println(searchInsert(nums1, 2))
}
