package main //https://leetcode.com/problems/count-odd-numbers-in-an-interval-range

func countOdds(low int, high int) int {
	noNum := (high - low + 1)
	if noNum%2 == 0 {
		return noNum / 2
	} else {
		if low%2 == 0 {
			return noNum / 2
		} else {
			return noNum/2 + 1
		}
	}
}
