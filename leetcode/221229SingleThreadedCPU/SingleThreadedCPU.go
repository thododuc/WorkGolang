package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for k := 0; k <= len(s)/2-1; k++ {
		if s[k] != s[len(s)-1-k] {
			return false
		}
	}
	return true
}

func main() {
	var x = 10
	fmt.Println(isPalindrome(x))
}
