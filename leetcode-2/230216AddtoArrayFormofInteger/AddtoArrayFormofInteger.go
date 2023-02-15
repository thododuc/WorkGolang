package main //https://leetcode.com/problems/add-to-array-form-of-integer

import (
	"fmt"
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	for i := len(num) - 1; i >= 0; i-- {
		num[i], k = (num[i]+k)%10, (num[i]+k)/10
		if k == 0 {
			return num
		}
	}
	kString := strconv.Itoa(k)
	kArray := make([]int, len(kString))
	for k, v := range kString {
		kArray[k] = int(v - '0')
	}
	num = append(kArray, num...)
	return num
}

func main() {
	num := []int{1, 2}
	k := 3485
	fmt.Println(addToArrayForm(num, k))
}
