package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if s == t {
		return true
	}
	if s == "" {
		return true
	}
	m := make(map[rune]bool)
	t1 := ""
	for _, v := range s {
		m[v] = true
	}
	for _, v := range t {
		if m[v] == true {
			t1 += string(v)
		}
	}

	if t1 == "" {
		return false
	}

	i := 0
	for j := 0; ; {
		if s[i] == t1[j] {
			i++
			if i == len(s) {
				return true
			}
		}
		j++
		if j == len(t1) {
			return false
		}
	}
}

func main() {
	s := ""
	t := "c"
	fmt.Println(isSubsequence(s, t))
}
