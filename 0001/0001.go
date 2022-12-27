// find 0 from right
// check array
// loop
package main

import "fmt"

func canJump(nums []int) bool {

	var result = true
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return true
		} else if nums[i] == 0 {
			result = false
			if i == 0 {
				return result
			}
			need := 1
			for j := i - 1; j >= 0; j-- {
				if nums[j] > need {
					result = true
					break
				}
				if j == 0 {
					return result
				}
				need++
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 0, 0}
	nums1 := []int{0}
	nums2 := []int{0, 1}
	nums3 := []int{2, 3, 1, 1, 4} //true
	nums4 := []int{3, 2, 1, 0, 4} //false
	nums5 := []int{2, 0, 1, 0, 1}
	fmt.Println(nums, ":", canJump(nums))
	fmt.Println(nums1, ":", canJump(nums1))
	fmt.Println(nums2, ":", canJump(nums2))
	fmt.Println(nums3, ":", canJump(nums3))
	fmt.Println(nums4, ":", canJump(nums4))
	fmt.Println(nums5, ":", canJump(nums5))
}
