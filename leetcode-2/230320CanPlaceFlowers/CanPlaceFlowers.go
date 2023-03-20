package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	justPut := false
	for i := 0; i < len(flowerbed); i++ {

		if flowerbed[i] == 0 {
			if !justPut {
				n--
				justPut = true
			} else {
				justPut = false
			}
		} else {
			if justPut {
				n++
			} else {
				justPut = true
			}
		}
	}
	return n <= 0
}
func main() {
	flowerbed := []int{0,1,0}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}