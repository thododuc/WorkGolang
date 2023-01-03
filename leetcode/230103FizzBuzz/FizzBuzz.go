package main //https://leetcode.com/problems/fizz-buzz/description/

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string { //khai bao chinh xac kich thuoc cua slice se giup tiet kiem bo nho
	result := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			result[i] = "FizzBuzz"
		} else if (i+1)%5 == 0 {
			result[i] = "Buzz"
		} else if (i+1)%3 == 0 {
			result[i] = "Fizz"
		} else {
			result[i] = strconv.Itoa(i + 1)
		}
	}
	return result
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
