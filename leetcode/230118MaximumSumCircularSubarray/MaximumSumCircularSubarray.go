package main //https://leetcode.com/problems/maximum-sum-circular-subarray/description/
import (
	"fmt"
	"math"
)

func maxSubarraySumCircular(nums []int) int {
	cur_max, end_max, cur_min, end_min, sum := 0, math.MinInt, 0, math.MaxInt, 0
	for _, v := range nums {
		cur_max += v
		if cur_max > end_max {
			end_max = cur_max
		}
		if cur_max < 0 {
			cur_max = 0
		}
		cur_min += v
		if cur_min < end_min {
			end_min = cur_min
		}
		if cur_min > 0 {
			cur_min = 0
		}
		sum += v
	}
	if end_max < sum-end_min && sum-end_min != 0 {
		end_max = sum - end_min
	}
	return end_max
}
func main() {
	nums := []int{-3, -2, -3}
	fmt.Println(maxSubarraySumCircular(nums))
}
