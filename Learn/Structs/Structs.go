package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(input string) *person {
	p := person{name: input}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(newPerson("Jon"))
}
