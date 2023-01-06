package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	r := [][]int{{0}}
	t := [][]int{}
	for i := 0; i < len(graph)-1; i++ {
		for j := 0; j < len(r); j++ {
			if r[j][len(r[j])-1] != i {
				t = append(t, r[j])
			} else {
				for z := 0; z < len(graph[i]); z++ {
					t = append(t, r[j])
					t[len(t)-1] = append(t[len(t)-1], graph[i][z])
				}
			}
		}
		r = t
		t = [][]int{}
	}
	return r
}

func main() {
	// graph := [][]int{{1, 2}, {3}, {3}, {}}
	// fmt.Println(allPathsSourceTarget(graph))
	// graph1 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	// fmt.Println(allPathsSourceTarget(graph1))
	graph2 := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(graph2))
}
