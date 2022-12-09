package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("slice s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("Slice:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	l = s[:]
	fmt.Println("sl2:", l)

	t := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlength := i + 1
		t[i] = make([]int, innerlength)
		for j := 0; j < innerlength; j++ {
			t[i][j] = i + j
		}
	}
	fmt.Println(t)
}
