package main

import "fmt"

func main() {
	arr := []int{1, 5, 5, 78, 9, 6, 3}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIdx := findPivot(arr, low, high)

		quickSort(arr, 0, pivotIdx - 1)
		quickSort(arr, pivotIdx + 1, high)
	}
}

func findPivot(arr []int, low, high int) int {
	i:= low -1

		for j:= low; j<high;j++{
			if arr[j] < arr[high] {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

	arr[i + 1], arr[high] = arr[high], arr[i+1]
	return i + 1
}