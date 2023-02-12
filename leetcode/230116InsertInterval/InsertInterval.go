package main //https://leetcode.com/problems/insert-interval

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	last := len(intervals) - 1
	a, b := -1, len(intervals)
	if intervals[0][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}
	if intervals[last][1] < newInterval[0] {
		return append(intervals, newInterval)
	}
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			a = i
		} else {
			break
		}
	}
	for ; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[1] {
			b = i
			break
		}
	}
	if intervals[a+1][0] < newInterval[0] {
		newInterval[0] = intervals[a+1][0]
	}
	if intervals[b-1][1] > newInterval[1] {
		newInterval[1] = intervals[b-1][1]
	}
	intervals = append(intervals[:a+2], intervals[b:]...)
	intervals[a+1] = newInterval
	return intervals
}

func main() {
	// intervals := [][]int{{1, 3}, {6, 9}}
	// newInterval := []int{2, 5}
	// fmt.Println(insert(intervals, newInterval))
	intervals1 := [][]int{{6, 10}, {13, 16}, {19, 19}, {23, 25}, {34, 39}, {41, 43}, {49, 51}}
	newInterval1 := []int{27, 27}
	fmt.Println(insert(intervals1, newInterval1))
}
