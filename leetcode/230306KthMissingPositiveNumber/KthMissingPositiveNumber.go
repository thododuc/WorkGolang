package main //https://leetcode.com/problems/kth-missing-positive-number

import "fmt"

func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}

	if arr[len(arr)-1 ] - len(arr) < k {
		return k + len(arr)
	}
	count, search := arr[0] - 1, arr[0] + 1
	for i := 1; i < len(arr);{
		if search == arr[i] {
			i++
		} else {
			count++
			if count == k {
				return search
			}
		}
		search++
	}
	return 0
}

func main() {
	arr := []int{1,2}
	 k := 1
	fmt.Println(findKthPositive(arr, k))
}