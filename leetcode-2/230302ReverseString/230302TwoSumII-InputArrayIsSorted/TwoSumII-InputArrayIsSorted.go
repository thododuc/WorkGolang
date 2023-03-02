package main //https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

import "fmt"

func twoSum(numbers []int, target int) []int {
	myMap := make(map[int]int)
	result := [2]int{}
	for k,v := range numbers{
		_, ok := myMap[v]
		if !ok {
			myMap[target - v] = k + 1
		} else {
			result[0] = myMap[v]
			result[1] = k + 1
			break
		}
		
	}
	return result[:]
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}