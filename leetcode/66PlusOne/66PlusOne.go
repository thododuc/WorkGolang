//https://leetcode.com/problems/plus-one/description/

package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] < 9 {
			digits[i]++
			return (digits)
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return (digits)
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	digits1 := []int{4, 3, 2, 2}
	fmt.Println(plusOne(digits1))
	digits2 := []int{8, 9, 9}
	fmt.Println(plusOne(digits2))
	digits3 := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits3))
}
