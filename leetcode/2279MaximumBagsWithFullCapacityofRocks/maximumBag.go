package main

import "fmt"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	cl := make([]int, 0)
	sum := 0
	for i := 0; i < len(capacity); i++ {
		var newNum = capacity[i] - rocks[i]
		if sum+newNum <= additionalRocks {
			if len(cl) == 0 {
				cl = append(cl, newNum)
			} else {
				if newNum >= cl[len(cl)-1] {
					cl = append(cl, newNum)
				} else {
					for j := 0; j < len(cl); j++ {
						if newNum <= cl[j] {
							cl = append(cl[:j+1], cl[j:]...)
							cl[j] = newNum
							break
						}
					}
				}
			}
			sum += newNum
		} else {
			if len(cl) != 0 && newNum < cl[len(cl)-1] {
				sum = sum + newNum - cl[len(cl)-1]
				for j := 0; j < len(cl); j++ {
					if newNum <= cl[j] {
						cl = append(cl[:j+1], cl[j:len(cl)-1]...)
						cl[j] = newNum
						break
					}
				}
			}
		}
	}
	return len(cl)
}

func main() {
	var additionalRocks = 77
	var capacity = []int{54, 18, 91, 49, 51, 45, 58, 54, 47, 91, 90, 20, 85, 20, 90, 49, 10, 84, 59, 29, 40, 9, 100, 1, 64, 71, 30, 46, 91}
	var rocks = []int{14, 13, 16, 44, 8, 20, 51, 15, 46, 76, 51, 20, 77, 13, 14, 35, 6, 34, 34, 13, 3, 8, 1, 1, 61, 5, 2, 15, 18}
	fmt.Println(maximumBags(capacity, rocks, additionalRocks))
	fmt.Println(len(capacity))
}
