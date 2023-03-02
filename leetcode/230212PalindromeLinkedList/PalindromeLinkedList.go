package main //https://leetcode.com/problems/palindrome-linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	valArray := []int{}
	for ; head != nil; head = head.Next {
		valArray = append(valArray, head.Val)
	}
	for i := 0; i < len(valArray)/2+1; i++ {
		if valArray[i] != valArray[len(valArray)-1-i] {
			return false
		}
	}
	return true
}
