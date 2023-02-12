package main //https://leetcode.com/problems/fruit-into-baskets

import "fmt"

func totalFruit(fruits []int) int {
	if len(fruits) <= 1 {
		return len(fruits)
	}
	t1, g1 := -1, 0
	t2, g2, gl, max := fruits[0], 1, 1, 1
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == t2 {
			g2++
			gl++
		} else if t1 == -1 {
			t1, t2 = t2, fruits[i]
			g1, g2, gl = g2, 1, 1
		} else if t1 != -1 {
			if fruits[i] == t1 {
				t1, t2 = t2, t1
				g1, g2, gl = g2, g1+1, 1
			} else {
				t1, t2 = t2, fruits[i]
				g1, g2, gl = gl, 1, 1
			}
		}
		if g1+g2 > max {
			max = g1 + g2
		}
	}
	return max
}

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(fruits))
}
