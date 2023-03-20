package main //https://leetcode.com/problems/pascals-triangle/

import "fmt"

func generate(numRows int) [][]int {
	result :=[][]int{{1}}
	if numRows == 1 {
		return result
	}
	for i:=2; i<=numRows; i++ {
		lastRow := make([]int,i)
		lastRow[0], lastRow[i-1] = 1, 1
		for j:=1; j <= i-2; j++ {
			lastRow[j] = result[i-2][j-1] + result[i-2][j]
		}
		result = append(result, lastRow)
	}
	return result
}

func main() {
	numRows := 10
	fmt.Println(generate(numRows))
}