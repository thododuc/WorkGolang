package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if (i+1)%5 == 0 {
			result = append(result, "Buzz")
		} else if (i+1)%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	return result
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
