package main

import "fmt"

func main() {
	a := [5]int{}
	b := make([]string, 5)
	c := make(map[string]int)

	c["k1"] = 7
	c["k2"] = 13

	delete(c, "k2")
	_, prs := c["k2"]
	fmt.Println(a, b)
	fmt.Println(prs)

}
