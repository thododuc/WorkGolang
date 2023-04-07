package main //https://leetcode.com/problems/pascals-triangle-ii

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	
	result := make([]int,rowIndex+1)
	result[0], result[1] = 1,1

	for i:=2;i <= rowIndex; i++ {
		for j:=i-1; j>0; j-- {
			result[j]+= result[j-1]
		}
		result[i] = 1
	}
	result[rowIndex] = 1
return result[:]
}

func main() {
	fmt.Println(getRow(5))
}