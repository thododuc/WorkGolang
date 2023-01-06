// find 0 from right
// check array
// loop
package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	lastIdx := len(nums) - 1

	record := make(map[int]bool)

	for i := lastIdx - 1; i > -1; i-- {
		if nums[i] >= lastIdx-i {
			record[i] = true
		} else {
			for key, _ := range record {
				if nums[i] >= key-i {
					record[i] = true
					break
				}
			}
		}
	}

	return record[0]
}

func main() {
	nums0 := []int{2, 3, 1, 1, 4}
	// nums := []int{2, 0, 0}
	// nums1 := []int{0}
	// nums2 := []int{0, 1}
	// nums3 := []int{2, 3, 1, 1, 4} //true
	// nums4 := []int{3, 2, 1, 0, 4} //false
	// nums5 := []int{2, 0, 1, 0, 1}
	fmt.Println(nums0, ":", canJump(nums0))
	// fmt.Println(nums, ":", canJump(nums))
	// fmt.Println(nums1, ":", canJump(nums1))
	// fmt.Println(nums2, ":", canJump(nums2))
	// fmt.Println(nums3, ":", canJump(nums3))
	// fmt.Println(nums4, ":", canJump(nums4))
	// fmt.Println(nums5, ":", canJump(nums5))
}
