package main

import "fmt"

func main() {
	t := make([][]int, 3)
	for i:= 0; i < 3; i++ {
		innerlength := i + 1
		t[] = make([]int, innerlength)
		for j:= 0; j < innerlength; j++ {
			t[i][j] = i + j
		}
	}
	fmt.Println(t)
}