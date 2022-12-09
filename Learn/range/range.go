package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 89}
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println(sum)

	myMap := map[string]string{"a": "airbus", "b": "boeing"}
	for k, v := range myMap {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
