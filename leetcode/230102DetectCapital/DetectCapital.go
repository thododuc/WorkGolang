// https://leetcode.com/problems/detect-capital/description/
package main

import "fmt"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	last := word[len(word)-1]
	if last >= 65 && last <= 90 {
		for _, v := range word[:len(word)-1] {
			if v < 65 || v > 90 {
				return false
			}
		}
	}
	if last >= 97 && last <= 122 {
		for _, v := range word[1 : len(word)-1] {
			if v < 97 || v > 122 {
				return false
			}
		}
	}
	return true
}

func main() {
	word := "a"
	fmt.Println(detectCapitalUse(word))
	word1 := "aaa"
	fmt.Println(detectCapitalUse(word1))
	word2 := "AA"
	fmt.Println(detectCapitalUse(word2))
	word3 := "aAa"
	fmt.Println(detectCapitalUse(word3))
	word4 := "aaaaB"
	fmt.Println(detectCapitalUse(word4))
}
