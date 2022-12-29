package main

import (
	"fmt"
	"sort"
)

func minStoneSum(piles []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(piles)))
	for i := 0; i < k; i++ {
		piles[0] -= piles[0] / 2
		heapify(piles, len(piles), 0)
	}
	var sum int
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}
	return sum
}

func heapify(arr []int, n int, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		temp := 0
		temp = arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp
		heapify(arr, n, largest)
	}
}

func main() {
	var piles = []int{5, 4, 9}
	var k = 2
	fmt.Println(minStoneSum(piles, k))
	var piles1 = []int{4, 3, 6, 7}
	var k1 = 3
	fmt.Println(minStoneSum(piles1, k1))
	var piles2 = []int{1, 3, 5, 7}
	var k2 = 3
	fmt.Println(minStoneSum(piles2, k2))
}
