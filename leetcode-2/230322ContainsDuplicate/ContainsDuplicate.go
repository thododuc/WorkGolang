package main //https://leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	myMap := make(map[int]bool)
	for _, v := range nums {
		if !myMap[v] {
			myMap[v] = true
		} else {
			return true
		}
	}
	return false
}

func main() {

}