package main //https://leetcode.com/problems/number-of-islands

import "fmt"

func dfs(r int, c int, grid [][]byte) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != 49 {
		return
	}
	grid[r][c] = 48
	dfs(r-1, c, grid)
	dfs(r, c+1, grid)
	dfs(r+1, c, grid)
	dfs(r, c-1, grid)
}
func numIslands(grid [][]byte) int {
	island := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 49 {
				island += 1
				dfs(i, j, grid)
			}
		}
	}
	return island
}

func main() {
	// grid := [][]byte{{49, 49, 48, 48, 48}, {49, 49, 48, 48, 48}, {48, 48, 49, 48, 48}, {48, 48, 48, 49, 49}}
	// fmt.Println(numIslands(grid))
	grid1 := [][]byte{{49}, {49}}
	fmt.Println(numIslands(grid1))
}
