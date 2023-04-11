package main //https://leetcode.com/problems/single-number/description/

func singleNumber(nums []int) int {
	myMap := make(map[int]bool)
	for _, v := range nums {
		myMap[v] = !myMap[v]
	}
	for k, _ := range myMap {
		if myMap[k] {
			return k
		}
	}
	return 0
}

func main() {

}