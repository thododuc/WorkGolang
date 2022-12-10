package main

import "fmt"

func intSeq() func() {
	a := 0
	b := 100
	return func() {
		a++
		b++
		if a%3 == 0 {
			fmt.Println(b)
		}
	}
}

func main() {
	testFunc := intSeq()
	testFunc()
	testFunc()
	testFunc()
	testFunc()
}
