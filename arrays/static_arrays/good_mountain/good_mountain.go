package main

import (
	"fmt"
)

func main() {
	var array = [7]int{}
	maxInt := -101
	maxIndex := 0
	for i := 0; i < 7; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 7; i++ {
		if array[i] > maxInt {
			maxInt = array[i]
			maxIndex = i
		}
	}
	if decreaseToMin(array, maxIndex) == raiseToMax(array, maxIndex) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
func raiseToMax(array [7]int, step1 int) bool {
	var flag bool
	for i := 0; i < step1; i++ {
		if array[i] < array[i+1] {
			flag = true
		} else {
			return false
		}
	}
	return flag
}
func decreaseToMin(array [7]int, step1 int) bool {
	var flag bool
	for i := step1; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			flag = true
		} else {
			return false
		}
	}
	return flag
}
