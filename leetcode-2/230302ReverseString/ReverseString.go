package main //https://leetcode.com/problems/reverse-string

import "fmt"

// func reverseString(s []byte) {
// 	for i := 0; i <= len(s)/2-1; i++ {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}