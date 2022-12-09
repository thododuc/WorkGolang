package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	for _, num := range mySlice {
		sum += num
	}
	fmt.Println(sum)
	for i, num := range mySlice {
		if num == 3 {
			fmt.Println(i)
		}
	}

	myMap := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range myMap {
		fmt.Println("%s -> %s\n", k, v)
	}
}
