//https://leetcode.com/problems/word-pattern/description/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	s1 := strings.Fields(s)

	m := make(map[rune]string)
	mapS := make(map[string]bool)

	if len(pattern) != len(s1) {
		return false
	}

	for k, v := range pattern {
		if m[v] == "" {
			if mapS[s1[k]] {
				return false
			}
			m[v] = s1[k]
			mapS[s1[k]] = true
		} else {
			if m[v] != s1[k] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "abb"
	s := "aa bb bb"
	fmt.Println(wordPattern(pattern, s))
}
