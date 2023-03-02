package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = string(sum%2+'0') + result
		carry = sum / 2
	}
	return result
}

func main() {
	a := "101"
	b := "110"
	fmt.Println(addBinary(a, b))
}
