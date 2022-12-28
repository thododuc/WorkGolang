func heapify(arr []int, n int, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		temp := 0
		temp = arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp
		heapify(arr, n, largest)
	}
}