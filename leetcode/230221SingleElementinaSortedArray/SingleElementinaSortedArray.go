package main //https://leetcode.com/problems/single-element-in-a-sorted-array

import "fmt"

// func singleNonDuplicate(nums []int) int {

// 	for i := 0; i < len(nums) && i+1 <len(nums); i += 2 {
// 		if nums[i] != nums[i+1] {
// 			return nums[i]
// 		}
// 	}
// 	return nums[len(nums)-1]
// }

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	if nums[left] != nums[left+1] {
		return nums[left]
	}
	if nums[right] != nums[right-1] {
		return nums[right]
	}
	mid := (left + right) / 2
	for left < right {
		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		} else {
			if nums[mid] == nums[mid-1] {
				mid -= 1
			}
			if (right - mid) % 2 != 0 {
				right = mid + 1
			} else {
				left = mid
			}
			mid =(left+right)/2
		}
	}
	return -1
}


func main() {
	// nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	// fmt.Println(singleNonDuplicate(nums))
	nums := []int{3,7,7,10,10,11,11}
	fmt.Println(singleNonDuplicate(nums))
}