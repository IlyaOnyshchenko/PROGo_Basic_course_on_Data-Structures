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
	balls := 0
	for i := 0; i < 10; i++ {
		nums := 1
		maxNums := 1
		currentIndex := 0
		for j := 1; j < len(array1); j++ {
			if array1[j] == array1[j-1] {
				nums++
			} else {
				if (nums > maxNums) && (nums > 2) {
					maxNums = nums
					balls += maxNums
					array1 = slices.Delete(array1, currentIndex, j)
					nums = 1
				} else {
					nums = 1
				}
			}
			if nums > 2 && j == len(array1)-1 {
				maxNums = nums
				balls += maxNums
				array1 = slices.Delete(array1, currentIndex, j+1)
			}
			if nums == 2 {
				currentIndex = j - 1
			}
		}
	}
	fmt.Println(balls)
}
