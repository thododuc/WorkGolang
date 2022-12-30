package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for ; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement((nums), val))
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val1 := 2
	fmt.Println(removeElement((nums1), val1))

}
