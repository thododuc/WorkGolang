package main //https://leetcode.com/problems/first-bad-version

import "fmt"

func isBadVersion(version int) bool {
	if version == 1 {
		return true
	} else {
		return false
	}
}

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	max, min, r := n, 1, 1
	for min <= max {
		i := (max + min) / 2
		if isBadVersion(i) {
			max = i - 1
			r = i
		} else {
			min = i + 1
		}

	}
	return r
}

func main() {
	n := 9
	fmt.Println(firstBadVersion(n))
}
