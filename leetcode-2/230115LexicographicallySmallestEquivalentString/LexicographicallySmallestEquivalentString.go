package main //https://leetcode.com/problems/lexicographically-smallest-equivalent-string

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	myMap := make(map[byte]string)
	result := ""
	for i := 0; i < len(s1); i++ {
		if myMap[s1[i]] == "" {
			myMap[s1[i]] = string(s1[i])
		}
		if myMap[s2[i]] == "" {
			myMap[s2[i]] = string(s2[i])
		}
		if myMap[s1[i]] != myMap[s2[i]] {
			temp := SortString(myMap[s1[i]] + myMap[s2[i]])
			for j := 0; j < len(myMap[s1[i]]); j++ {
				myMap[myMap[s1[i]][j]] = temp
			}
			for j := 0; j < len(myMap[s2[i]]); j++ {
				myMap[myMap[s2[i]][j]] = temp
			}
		}
	}
	fmt.Println(myMap)
	for k, v := range myMap {
		myMap[k] = string(v[0])
	}
	fmt.Println(myMap)
	for i := 0; i < len(baseStr); i++ {
		if myMap[baseStr[i]] == "" {
			myMap[baseStr[i]] = string(baseStr[i])
		}
		result += myMap[baseStr[i]]
	}
	return result
}

func main() {
	s1 := "leetcode"
	s2 := "programs"
	baseStr := "sourcecode"
	fmt.Println(smallestEquivalentString(s1, s2, baseStr))
}
