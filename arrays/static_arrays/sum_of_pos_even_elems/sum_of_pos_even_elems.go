package main

import (
	"fmt"
)

func main() {
	var array = [5]int{}
	var sum int
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 5; i++ {
		if array[i]%2 == 0 && array[i] > 0 {
			sum += array[i]
		}
	}
	fmt.Println(sum)
}
