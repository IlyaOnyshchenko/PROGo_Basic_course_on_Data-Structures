package main

import (
	"fmt"
)

func main() {
	var sizeArray1, sizeArray2 int
	fmt.Scan(&sizeArray1)
	array1 := make([]int, sizeArray1)
	for i := 0; i < sizeArray1; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Scan(&sizeArray2)
	array2 := make([]int, sizeArray2)
	for i := 0; i < sizeArray2; i++ {
		fmt.Scan(&array2[i])
	}
	for i := 0; i < sizeArray2; i++ {
		array1 = append(array1, array2[i])
	}
	counterSort(array1, len(array1))
}
func counterSort(list []int, width int) {
	min := 100001
	max := -100001
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}
	counts := make([]int, max+intAbs(min)+1)
	for i := 0; i < width; i++ {
		counts[list[i]+intAbs(min)] += 1
	}
	for i := 0; i < (max+intAbs(min))+1; i++ {
		for j := 0; j < counts[i]; j++ {
			if counts[i] != 0 {
				fmt.Print(i-intAbs(min), " ")
			}
		}
	}
}
func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
