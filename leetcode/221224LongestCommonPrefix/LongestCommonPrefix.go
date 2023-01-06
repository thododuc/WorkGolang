package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var prefix string
	prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		if prefix == "" {
			break
		}
		j := 0
		for ; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				break
			}
		}
		prefix = prefix[0:j]
	}
	return prefix
}

func main() {
	strs := []string{"acc", "aaa", "aaba"}
	fmt.Println(longestCommonPrefix(strs))
}
