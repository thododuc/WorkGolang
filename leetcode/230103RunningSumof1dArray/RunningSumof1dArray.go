package main

import "fmt"

func runningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}
