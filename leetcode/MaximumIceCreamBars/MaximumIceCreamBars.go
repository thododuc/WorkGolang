package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for k, v := range costs {
		coins -= v
		if coins < 0 {
			return k
		}
	}
	return len(costs)
}

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7
	fmt.Println(maxIceCream(costs, coins))
}
