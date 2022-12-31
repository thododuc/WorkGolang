package main

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	var num1, num2, sum = new(big.Int), new(big.Int), new(big.Int) //num1, num2, sum la pointer *big.Int mang gia tri 0
	num1.SetString(a, 2)
	num2.SetString(b, 2)
	sum.Add(num1, num2)
	return (sum.Text(2))
}

func main() {
	a := "11001"
	b := "11"
	fmt.Println(addBinary(a, b))
}
