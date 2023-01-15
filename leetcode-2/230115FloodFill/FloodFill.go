package main //https://leetcode.com/problems/flood-fill

import "fmt"

func urbf(image [][]int, sr int, sc int, begin int, color int) [][]int {
	around := [][]int{}
	if sr-1 >= 0 && image[sr-1][sc] == begin {
		image[sr-1][sc] = color
		around = append(around, []int{sr - 1, sc})
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] == begin {
		image[sr][sc+1] = color
		around = append(around, []int{sr, sc + 1})
	}
	if sr+1 < len(image) && image[sr+1][sc] == begin {
		image[sr+1][sc] = color
		around = append(around, []int{sr + 1, sc})
	}
	if sc-1 >= 0 && image[sr][sc-1] == begin {
		image[sr][sc-1] = color
		around = append(around, []int{sr, sc - 1})
	}
	return around
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	begin := image[sr][sc]
	field := [][]int{{sr, sc}}
	image[sr][sc] = color
	for len(field) != 0 {
		field = append(field[1:], urbf(image, field[0][0], field[0][1], begin, color)...)
	}
	return image
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, color := 1, 1, 2
	fmt.Println(floodFill(image, sr, sc, color))
}
