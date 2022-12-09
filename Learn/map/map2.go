package main

import "fmt"

func main() {
	var myMap = make(map[float64]string)
	myMap[1.3] = "Batman"
	myMap[1.5] = "Robin"
	fmt.Println(myMap)
	map2 := map[int]string{
		90: "dog",
		91: "cat",
		92: "wolf",
		93: "tiger",
	}

	map2[100] = "dragon"
	for id, pet := range map2 {
		fmt.Println(id, pet)
	}
	fmt.Println(map2[95])
	petname, ok1 := map2[92]
	fmt.Println("\nKey present or not:", ok1)
	fmt.Println("Pet no 92:", petname)
	map3 := map2

	map3[96] = "bird"
	fmt.Println("map2 after modify map3:", map2)

}
