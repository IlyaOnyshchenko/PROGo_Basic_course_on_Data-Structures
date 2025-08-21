package main

import (
	"fmt"
	"slices"
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
	var ind int
	array3 := []int{}
	for i := 0; i < len(array1); i++ {
		if slices.Contains(array2, array1[i]) == false {
			array3 = slices.Insert(array3, ind, array1[i])
			ind++
		}
	}
	for _, value := range array3 {
		fmt.Print(value, " ")
	}
}
