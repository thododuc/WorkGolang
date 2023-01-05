package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	arrow := 1
	first := points[0][1]

	for i := 0; i < len(points); i++ {
		if points[i][0] > first {
			first = points[i][1]
			arrow++
		}
	}
	return arrow
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
	points1 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(findMinArrowShots(points1))
	points2 := [][]int{{2, 3}, {2, 3}}
	fmt.Println(findMinArrowShots(points2))
	points3 := [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}
	fmt.Println(findMinArrowShots(points3))
}
