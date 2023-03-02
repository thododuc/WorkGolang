package main //https://leetcode.com/problems/rotate-array

import "fmt"

func rotate(nums []int, k int) {
	if k >=len(nums) {
			k = k%len(nums)
	}
	fmt.Println("k:",k)
temp :=[]int{}
for i:=len(nums)-k; i<len(nums); i++ {
	temp=append(temp, nums[i])
}

for i:=len(nums)-1; i>= k; i-- {
	nums[i] = nums[i-k]
}
for i:=0; i < k; i++ {
	nums[i] = temp[i]
}
}

func main() {
	nums := []int{1, 2}
	rotate(nums, 5)
	fmt.Println(nums)
}