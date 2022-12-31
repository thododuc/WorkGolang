package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits); ; {
		i--
		if i == -1 {
			digits = append([]int{1}, digits...)
			return (digits)
		}
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return (digits)
		}
	}
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
