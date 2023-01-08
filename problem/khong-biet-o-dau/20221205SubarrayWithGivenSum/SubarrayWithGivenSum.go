package main

import (
	"fmt"
	"time"
)

func main() {
	start_time := time.Now()
	// n := 5
	s := 15
	a := [74]int{142, 112, 54, 69, 148, 45, 63, 158, 38, 60, 124, 142, 130, 179, 117, 36, 191, 43, 89, 107, 41, 143, 65, 49, 47, 6, 91, 130, 171, 151, 7, 102, 194, 149, 30, 24, 85, 155, 157, 41, 167, 177, 132, 109, 145, 40, 27, 124, 138, 139, 119, 83, 130, 142, 34, 116, 40, 59, 105, 131, 178, 107, 74, 187, 22, 146, 125, 73, 71, 30, 178, 174, 98, 113}
	var low, hight int
	var sum int
	low = 0
	hight = 0
	if s == 0 {
		fmt.Println(-1)
		return
	}
	for hight = 0; hight < len(a); {
		sum += a[hight]
		if sum == s {
			fmt.Println(low+1, hight+1)
			return
		}
		if sum > s {
			sum -= a[low]
			low++
		}
		hight++
	}
	fmt.Println(-1)
	time_exe := time.Since(start_time)
	fmt.Println(time_exe.Microseconds())
}
