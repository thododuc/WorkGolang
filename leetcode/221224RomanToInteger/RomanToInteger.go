package main

import "fmt"

func romanToInt(s string) int {
	var sum int
	var lastNum, currNum int
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			currNum = 1
		case 'V':
			currNum = 5
		case 'X':
			currNum = 10
		case 'L':
			currNum = 50
		case 'C':
			currNum = 100
		case 'D':
			currNum = 500
		case 'M':
			currNum = 1000
		}
		if currNum >= lastNum {
			sum += currNum
		} else {
			sum -= currNum
		}
		lastNum = currNum
	}

	return (sum)
}

func main() {
	var s string
	// s = "LXXXIX"
	fmt.Println("Enter Roman number:")
	fmt.Scanf("%s", &s)
	fmt.Println(romanToInt(s))
}
