package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sumLeft := 0
	for k, v := range nums {
		if sumLeft*2 == sum-v {
			return k
		} else {
			sumLeft += v
		}
	}
	return -1
}

func main() {
	nums := []int{-1, -1, -1, -1, -1, -1}
	fmt.Println(pivotIndex(nums))
}
