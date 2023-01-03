package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	ransomNote1 := strings.Split(ransomNote, "")
	for _, v := range ransomNote1 {
		idx := strings.Index(magazine, v)
		if idx == -1 {
			return false
		} else {
			if idx == len(magazine) {
				magazine = magazine[:idx]
			} else {
				magazine = magazine[:idx] + magazine[idx+1:]
			}
		}
	}
	return true
}

func main() {
	ransomNote := "aab"
	magazine := "baa"
	fmt.Println(canConstruct(ransomNote, magazine))
}
