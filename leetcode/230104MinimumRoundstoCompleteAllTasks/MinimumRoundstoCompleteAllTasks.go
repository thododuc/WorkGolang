package main

import "fmt"

func minimumRounds(tasks []int) int {
	r := 0
	m := make(map[int]int)
	for _, v := range tasks {
		m[v]++
	}
	for _, v := range m {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			r += v / 3
		} else {
			r += v/3 + 1
		}
	}
	return r
}

func main() {
	tasks := []int{2, 3, 3}
	fmt.Println(minimumRounds(tasks))
}
