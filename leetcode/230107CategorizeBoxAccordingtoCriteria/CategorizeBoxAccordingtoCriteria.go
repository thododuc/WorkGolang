package main

import (
	"fmt"
	"math"
)

func categorizeBox(length int, width int, height int, mass int) string {
	var bulky, heavy bool
	l, w, h := uint32(length), uint32(width), uint32(height)
	if l >= uint32(math.Pow(10, 4)) || w >= uint32(math.Pow(10, 4)) || h >= uint32(math.Pow(10, 4)) || l*w*h >= uint32(math.Pow(10, 9)) {
		bulky = true
	}
	if mass >= 100 {
		heavy = true
	}

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
	length, width, height, mass := 1000, 35, 700, 300
	fmt.Println(categorizeBox(length, width, height, mass))
}
