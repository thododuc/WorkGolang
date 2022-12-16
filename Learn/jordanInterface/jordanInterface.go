package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	d := Dog{}
	// animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}} //animals la 1 array co type Animal. Tuong tu b := [5]int{1,2,3,4,5}
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }
	fmt.Println(d.Speak())
}
