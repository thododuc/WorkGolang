package main //https://leetcode.com/problems/single-element-in-a-sorted-array

import "fmt"

func singleNonDuplicate(nums []int) int {

	for i := 0; i < len(nums) && i+1 <len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}