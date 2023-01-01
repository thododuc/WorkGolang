//https://leetcode.com/problems/word-pattern/description/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	s1 := strings.Fields(s)

	m := make(map[rune]string)

	if len(pattern) != len(s1) {
		return false
	}

	for k, v := range pattern {
		if m[v] == "" {
			for _, mv := range m {
				if mv == s1[k] {
					return false
				}
			}
			m[v] = s1[k]
		} else {
			if m[v] != s1[k] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "aaaaa"
	s := "aa aa aa aa"
	fmt.Println(wordPattern(pattern, s))
}
