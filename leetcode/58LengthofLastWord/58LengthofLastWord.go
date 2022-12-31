package main

import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	t := 0
	for i >= 0 && s[i] == 32 {
		i--
	}
	for i >= 0 && s[i] != 32 {
		t++
		i--
	}
	return t
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
	s1 := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s1))
	s2 := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s2))
}
