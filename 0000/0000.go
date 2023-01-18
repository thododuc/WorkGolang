package main //https://leetcode.com/problems/flood-fill

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if m[v] != 0 {
			return []int{m[v], k}
		} else {
			m[target-v] = k
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
