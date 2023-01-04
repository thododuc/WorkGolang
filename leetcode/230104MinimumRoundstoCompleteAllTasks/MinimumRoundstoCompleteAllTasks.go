package main //https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/description/

import (
	"fmt"
	"sort"
)

// func minimumRounds(tasks []int) int {
// 	r := 0
// 	m := make(map[int]int)
// 	for _, v := range tasks {
// 		m[v]++
// 	}
// 	for _, v := range m {
// 		if v == 1 {
// 			return -1
// 		}
// 		if v%3 == 0 {
// 			r += v / 3
// 		} else {
// 			r += v/3 + 1
// 		}
// 	}
// 	return r
// }

func minimumRounds(tasks []int) int {
	if len(tasks) == 1 {
		return -1
	}
	sort.Ints(tasks)
	count := 1
	r := 0
	for i := 1; i < len(tasks); i++ {
		if tasks[i] > tasks[i-1] {
			if count == 1 {
				return -1
			}
			if count%3 == 0 {
				r += count / 3
			} else {
				r += count/3 + 1
			}
			count = 1
		} else {
			count++
		}
	}
	if count == 1 {
		return -1
	}
	if count%3 == 0 {
		r += count / 3
	} else {
		r += count/3 + 1
	}
	return r
}

func main() {
	tasks := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	fmt.Println(minimumRounds(tasks))
}
