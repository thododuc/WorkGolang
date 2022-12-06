package main

import "fmt"

func main() {
	a := [5]int{}
	b := make([]string, 3)
	c := make(map[string]int)

	c["k1"] = 7
	c["k2"] = 13

	fmt.Println(a, b)
	fmt.Println(c)

	var d map[string]int
	fmt.Println(d)
}
