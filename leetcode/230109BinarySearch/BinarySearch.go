package main

import "fmt"

func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	min, max := 0, len(nums)-1
	i := 0
	for min <= max {
		i = (min + max) / 2
		if nums[i] < target {
			min = i + 1
		} else if nums[i] > target {
			max = i - 1
		}
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 0
	fmt.Println(search(nums, target))
}
