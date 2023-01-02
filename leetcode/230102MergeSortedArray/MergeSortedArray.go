//https://leetcode.com/problems/merge-sorted-array/description/

package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 6
	num2 := []int{2, 5, 6}
	n := 3
	merge(num1, m, num2, n)
	fmt.Println(num1)
}
