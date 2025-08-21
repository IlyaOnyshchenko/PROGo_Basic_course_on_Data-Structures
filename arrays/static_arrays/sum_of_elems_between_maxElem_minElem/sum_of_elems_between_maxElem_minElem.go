package main

import (
	"fmt"
)

func main() {
	var array = [5]int{}
	maxInt := -101
	minInt := 101
	maxIndex := 0
	minIndex := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 5; i++ {
		if array[i] > maxInt {
			maxInt = array[i]
			maxIndex = i
		}
		if array[i] < minInt {
			minInt = array[i]
			minIndex = i
		}
	}
	if minIndex > maxIndex {
		fmt.Println(rightLeft(array, minIndex, maxIndex))
	} else {
		fmt.Println(leftRight(array, minIndex, maxIndex))
	}
}
func leftRight(array1 [5]int, step1 int, step2 int) int {
	var sum int
	for i := step1 + 1; i < step2; i++ {
		sum += array1[i]
	}
	return sum
}
func rightLeft(array2 [5]int, step1 int, step2 int) int {
	var sum int
	for i := step2 + 1; i < step1; i++ {
		sum += array2[i]
	}
	return sum
}
