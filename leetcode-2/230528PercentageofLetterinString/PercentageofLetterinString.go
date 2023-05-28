package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			num++
		}
	}
	return num * 100 / len(s)
}

func main() {
	fmt.Println(percentageLetter("foobar", 111))
}
