package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("array a:", a)
	a[4] = 1000
	fmt.Println("set a[4]:", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = j
		}
	}
	fmt.Println(c)
}
