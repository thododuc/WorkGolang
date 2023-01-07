package main //https://leetcode.com/problems/gas-station/description/

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}
	count := 0
	sum = 0
	for _, v := range gas {
		sum += v
		if sum >= 0 {
			count++
		} else {
			sum = 0
			count = 0
		}
	}
	return len(gas) - count
}

// func canCompleteCircuit(gas []int, cost []int) int {
// 	sum := 0
// 	station := 0
// 	tank := 0
// 	for i := 0; i < len(gas); i++ {
// 		sum += gas[i] - cost[i]
// 		tank += gas[i] - cost[i]
// 		fmt.Println(sum)
// 		if tank < 0 {
// 			station = i + 1
// 			tank = 0
// 		}
// 	}
// 	if sum < 0 {
// 		return -1
// 	}
// 	return station
// }

func main() {
	gas := []int{1, 2, 3, 4, 5, 5, 70}
	cost := []int{2, 3, 4, 3, 9, 6, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas1, cost1))
}
