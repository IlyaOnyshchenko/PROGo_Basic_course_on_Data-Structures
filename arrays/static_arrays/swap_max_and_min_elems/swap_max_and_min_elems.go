package main

import (
	"fmt"
)

func main() {
	var array = [4]int{}
	maxInt := -1001
	minInt := 1001
	maxIndex := 0
	minIndex := 0
	for i := 0; i < 4; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 4; i++ {
		if array[i] > maxInt {
			maxInt = array[i]
			maxIndex = i
		}
		if array[i] < minInt {
			minInt = array[i]
			minIndex = i
		}
	}
	array[maxIndex], array[minIndex] = array[minIndex], array[maxIndex]
	for _, number := range array {
		fmt.Print(number, " ")
	}
}
