// https://leetcode.com/problems/climbing-stairs/description/
package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	} else {
		a := 1
		b := 2
		c := 3
		for i := 3; i < n+1; i++ {
			c = a + b
			a = b
			b = c
		}
		return c
	}
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}
