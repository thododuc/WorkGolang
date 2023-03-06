package main //https://leetcode.com/problems/kth-missing-positive-number

import "fmt"

func findKthPositive(arr []int, k int) int {
	count, search := 0, 0

	for i := 0; i < len(arr);{
		search++
		if search == arr[i] {
			i++
		} else {
			count++
			if count == k {
				return search
			}
		}
	}
	return search + k - count
}

func main() {
	arr := []int{2,3,4,7,11}
	 k := 5
	fmt.Println(findKthPositive(arr, k))
}