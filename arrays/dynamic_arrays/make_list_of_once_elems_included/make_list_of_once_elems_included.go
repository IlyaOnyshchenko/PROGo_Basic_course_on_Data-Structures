package main

import (
	"fmt"
	"slices"
)

func main() {
	var sizeArray1 int
	fmt.Scan(&sizeArray1)
	array1 := make([]int, sizeArray1)
	for i := 0; i < sizeArray1; i++ {
		fmt.Scan(&array1[i])
	}
	var ind int
	array2 := []int{}
	for i := 0; i < len(array1); i++ {
		if Find(array1, array1[i]) == 1 {
			array2 = slices.Insert(array2, ind, array1[i])
			ind++
		}
	}
	for _, value := range array2 {
		fmt.Print(value, " ")
	}
}
func Find(array []int, key int) int {
	var quantity int
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			quantity++
		}
	}
	return quantity
}
