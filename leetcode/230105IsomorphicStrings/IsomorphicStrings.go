package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m1 := make(map[rune]byte)
	m2 := make(map[byte]bool)
	for k, v := range s {
		if m1[v] == 0 {
			if m2[t[k]] {
				return false
			} else {
				m1[v] = t[k]
				m2[t[k]] = true
			}
		} else {
			if m1[v] != t[k] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "foo"
	t := "bar"
	fmt.Println(isIsomorphic(s, t))
}
