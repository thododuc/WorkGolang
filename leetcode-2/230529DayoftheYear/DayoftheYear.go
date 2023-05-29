// https://leetcode.com/problems/day-of-the-year/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	dateConv := strings.Split(date, "-")
	year,_ := strconv.Atoi(dateConv[0])
	month,_ := strconv.Atoi(dateConv[1])
	day,_ := strconv.Atoi(dateConv[2])

	dayInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	
	if year % 4 ==0 && (year % 100 != 0 || year % 400 == 0) {
		dayInMonth[1] = 29
	}

	dayNumber := 0

	for i:= 0; i< month - 1 ; i++{
		dayNumber += dayInMonth[i]
	}

	dayNumber += day

	return dayNumber
}

func main() {
	date := "2023-10-09"
	fmt.Println(dayOfYear(date))
}