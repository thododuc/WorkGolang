package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	j := 1

	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	var nums1 = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums1))
	var nums2 = []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	fmt.Println(removeDuplicates(nums2))
}
