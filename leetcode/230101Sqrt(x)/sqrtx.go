package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	min := 2
	max := x / 2
	r := (min + max) / 2
	for min < max-1 {
		if r*r == x {
			return r
		} else if r*r < x {
			min = r
		} else {
			max = r
		}
		r = (min + max) / 2
	}
	return r
}

func main() {
	x := 15465871
	fmt.Println(mySqrt(x))
	fmt.Println(math.Sqrt(float64(x)))
}
