package main

import "fmt"

func isPalindrome(head *ListNode) bool {
	for i:=0; i < len(head)/2 ; i++ {
			if head[i] != head[n-1-i] {
				return false
			}
	}
	return true
}

func main() {
	var head := []int{1,2,2,1}
	fmt.Println(isPalindrome(head))
}