package main

import "fmt"

func main() {

	var fibo func(n int) int
	fibo = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fibo(n-1) + fibo(n-2)
		}
	}
	fmt.Println(fibo(7))
}
