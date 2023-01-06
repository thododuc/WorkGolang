package main

import (
	"fmt"
)

func main() {
	var i int
	var p *int
	var b bool
	b = i == *p
	b = i <= *p
	b = i >= *p
	i += *p
	i -= *p
	i *= *p
	i /= *p
	i %= *p
	i = *p << uint(i)
	i = *p >> uint(i)
	fmt.Println(i, p, b)
}
