package main

import "fmt"

func main() {
	slice1 := []int{1, 3, 4, 5}
	slice2 := []int{2, 4, 6, 8}
	// slice1 = insert(slice1, 1, slice2[2])
	slice1 = append(slice1[:2], slice1[1:]...)
	slice1[1] = slice2[2]

	fmt.Println(slice1) // [1 6 3 4 5]
}
