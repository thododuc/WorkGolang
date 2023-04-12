package main //https://leetcode.com/problems/arranging-coins/description/

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	sqrt_n := int(math.Sqrt(float64(n*2)))
	if n* 2 >= sqrt_n * (sqrt_n+1) {
		return sqrt_n +1
	} else {
		return sqrt_n
	}
}

func main() {
fmt.Println(arrangeCoins(8))
}