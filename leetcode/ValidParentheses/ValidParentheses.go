package main

import "fmt"

func isValid(s string) bool {
	c := true
	var p []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			p = append(p, 1)
		case '[':
			p = append(p, 2)
		case '{':
			p = append(p, 3)
		case ')':
			if len(p) == 0 {
				c = false
				return c
			} else if p[len(p)-1] == 1 {
				p = p[:len(p)-1]
			} else {
				c = false
				break
			}
		case ']':
			if len(p) == 0 {
				c = false
				return c
			} else if p[len(p)-1] == 2 {
				p = p[:len(p)-1]
			} else {
				c = false
				break
			}
		case '}':
			if len(p) == 0 {
				c = false
				return c
			} else if p[len(p)-1] == 3 {
				p = p[:len(p)-1]
			} else {
				c = false
				break
			}
		}
	}
	if len(p) != 0 {
		c = false
	}
	return c
}

func main() {
	s := "(){}}{"
	fmt.Println(isValid(s))
}
