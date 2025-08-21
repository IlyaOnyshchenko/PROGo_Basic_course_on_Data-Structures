package main

import (
	"fmt"
	"slices"
)

func main() {
	var size, elementToDelete int
	fmt.Scan(&size)
	array := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&elementToDelete)
	for slices.Contains(array, elementToDelete) == true {
		array = slices.Delete(array, slices.Index(array, elementToDelete), slices.Index(array, elementToDelete)+1)
	}
	for _, v := range array {
		fmt.Print(v, " ")
	}
}
