package main //

import (
	"fmt"
)

func isHappy(n int) bool {
	myMap := make(map[int]bool)
	for {
    myMap[n] = true
		sum := 0
		for n > 0 {
			bp := n%10
			sum = sum + bp * bp
			n /= 10
		}
		n = sum
		if myMap[n] {
			return false
		} else if n == 1 {
			return true
		}
	}
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}