package main

import (
	"fmt"
	"slices"
)

func main() {
	var number, position, element int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&position, &element)
	array = slices.Insert(array, position, element)
	for _, v := range array {
		fmt.Print(v, " ")
	}
}
