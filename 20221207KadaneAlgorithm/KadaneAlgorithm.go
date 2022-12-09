package main

import "fmt"

func main() {
	//n := 5
	arr := [5]int{1, 2, 3, -2, 5}
	var sum, max int
	max = arr[0]
	for i := range arr {
		sum += arr[i]
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Println(max)
}
