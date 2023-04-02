package main //https://leetcode.com/problems/contains-duplicate-ii

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		k1:= k
		for j := i + 1; j < len(nums) && k1 > 0; j++ {
			if nums[i] == nums[j] {
				return true
			}
			k1--
		}
	}
	return false
}

func main() {
	nums := []int{1,2,3,1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}