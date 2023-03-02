package main //https://leetcode.com/problems/move-zeroes

import "fmt"

func moveZeroes(nums []int) {
	numOfZero := 0
	temp := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if numOfZero == 0 {
				temp = i
			}
			numOfZero++
		} else {
			if numOfZero != 0 {
				nums[temp] = nums[i]
				nums[i]=0
				temp++
			}
		}
	}
}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}