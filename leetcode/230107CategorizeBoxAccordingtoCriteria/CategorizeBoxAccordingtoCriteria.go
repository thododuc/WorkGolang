package main //https://leetcode.com/contest/biweekly-contest-95/problems/categorize-box-according-to-criteria/

import (
	"fmt"
	"math"
)

func categorizeBox(length int, width int, height int, mass int) string {
	bulky, heavy := false, false
	l, w, h := uint32(length), uint32(width), uint32(height)
	if l >= uint32(math.Pow(10, 4)) || w >= uint32(math.Pow(10, 4)) || h >= uint32(math.Pow(10, 4)) || l*w*h >= uint32(math.Pow(10, 9)) {
		bulky = true
	}
	if mass >= 100 {
		heavy = true
	}
	fmt.Println(l * w * h)
	switch {
	case bulky && heavy:
		return "Both"
	case !bulky && !heavy:
		return "Neither"
	case !bulky && heavy:
		return "Heavy"
	default:
		return "Bulky"
	}
}

func main() {
	fmt.Println(math.Pow(10, 4), math.Pow(10, 9))
	length, width, height, mass := 1000, 35, 700, 300
	fmt.Println(categorizeBox(length, width, height, mass))
}
