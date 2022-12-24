package main

import "fmt"

func isValid(s string) bool {
	tr := map[rune]rune{
		'}': '{',
		'[': ']',
		'(': ')',
	}
	var p []rune
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			p = append(p, ch)
		case ')', ']', '}':
			if len(p) == 0 || p[len(p)-1] != tr[ch] {
				return false
			} else {
				p = p[:len(p)-1]
			}
		}
	}
	return len(p) == 0
}

func main() {
	s := "(){}}{"
	fmt.Println(isValid(s))
}

// Day la loi giai ban dau cua minh. Y tuong thi trung khop nhung code dai dong hon nhieu.
// func isValid(s string) bool {
// 	c := true
// 	var p []int
// 	for i := 0; i < len(s); i++ {
// 		switch s[i] {
// 		case '(':
// 			p = append(p, 1)
// 		case '[':
// 			p = append(p, 2)
// 		case '{':
// 			p = append(p, 3)
// 		case ')':
// 			if len(p) == 0 {
// 				c = false
// 				return c
// 			} else if p[len(p)-1] == 1 {
// 				p = p[:len(p)-1]
// 			} else {
// 				c = false
// 				break
// 			}
// 		case ']':
// 			if len(p) == 0 {
// 				c = false
// 				return c
// 			} else if p[len(p)-1] == 2 {
// 				p = p[:len(p)-1]
// 			} else {
// 				c = false
// 				break
// 			}
// 		case '}':
// 			if len(p) == 0 {
// 				c = false
// 				return c
// 			} else if p[len(p)-1] == 3 {
// 				p = p[:len(p)-1]
// 			} else {
// 				c = false
// 				break
// 			}
// 		}
// 	}
// 	if len(p) != 0 {
// 		c = false
// 	}
// 	return c
// }
