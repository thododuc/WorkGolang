package main //https://leetcode.com/problems/number-of-1-bits

func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		if num%2 != 0 {
			result++
		}
		num /= 2
	}
	return result
}

func main() {

}