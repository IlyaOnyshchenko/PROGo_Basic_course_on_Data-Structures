package main

import (
	"fmt"
	"slices"
)

func main() {
	var number, position int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&position)
	array = slices.Delete(array, position, position+1)
	for _, v := range array {
		fmt.Print(v, " ")
	}
}
