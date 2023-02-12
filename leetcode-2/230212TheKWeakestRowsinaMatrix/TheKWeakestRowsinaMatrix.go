package main //The K Weakest Rows in a Matrix

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	mat2 := [][]int{}
	for i := 0; i < len(mat); i++ {
		soldier := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				soldier++
			}
		}
		mat2 = append(mat2, []int{soldier, i})
	}
	fmt.Println(mat2)
	sort.SliceStable(mat2, func(i, j int) bool { return mat2[i][0] < mat2[j][0] })
	mat3 := []int{}
	for i := 0; i < k; i++ {
		mat3 = append(mat3, mat2[i][1])
	}

	return mat3
}

func main() {
	mat := [][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}
	k := 6
	fmt.Println(kWeakestRows(mat, k))
}
