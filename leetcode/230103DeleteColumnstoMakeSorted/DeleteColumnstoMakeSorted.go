package main

import "fmt"

func minDeletionSize(strs []string) int {
	result := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				result++
				break
			}
		}
	}
	return result
}

func main() {
	// strs := []string{"cba", "daf", "ghi"}
	// fmt.Println(minDeletionSize(strs))
	// strs1 := []string{"a", "b"}
	// fmt.Println(minDeletionSize(strs1))
	// strs2 := []string{"zyx", "wvu", "tsr"}
	// fmt.Println(minDeletionSize(strs2))
	strs3 := []string{"rrjk", "furt", "guzm"}
	fmt.Println(minDeletionSize(strs3))
}
