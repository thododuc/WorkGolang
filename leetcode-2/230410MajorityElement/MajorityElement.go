package main //https://leetcode.com/problems/majority-element/description/

func majorityElement(nums []int) int {
	myMap := make(map[int]int)
	majority := len(nums) / 2
	for _, v := range nums {
		myMap[v]++
		if myMap[v] > majority {
			return v
		}
	}
	return 0
}

func main() {

}