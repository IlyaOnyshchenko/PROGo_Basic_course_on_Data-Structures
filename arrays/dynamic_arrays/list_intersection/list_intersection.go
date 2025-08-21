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
	intersectSortedLists(array1, array2)
}
func intersectSortedLists(list1, list2 []int) {
	p1, p2 := 0, 0
	result := []int{}
	for p1 < len(list1) && p2 < len(list2) {
		if list1[p1] == list2[p2] {
			result = append(result, list1[p1])
			p1++
			p2++
		} else if list1[p1] < list2[p2] {
			p1++
		} else {
			p2++
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i], " ")
	}
}
