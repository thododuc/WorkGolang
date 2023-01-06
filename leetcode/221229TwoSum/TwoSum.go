package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, val := range nums {
		if mp[target-val] != 0 {
			return []int{mp[target-val] - 1, i}
		}
		mp[val] = i + 1
	}

	return []int{}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}
