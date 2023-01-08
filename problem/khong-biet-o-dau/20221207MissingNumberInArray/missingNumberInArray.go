package main

import "fmt"

func main() {
	n := 10
	arr := [9]int{6, 1, 2, 8, 3, 4, 7, 10, 5}
	var sumN, sumA int
	sumN = n * (n + 1) / 2
	for i := 0; i < 9; i++ {
		sumA += arr[i]
	}
	fmt.Println(sumN - sumA)
}
