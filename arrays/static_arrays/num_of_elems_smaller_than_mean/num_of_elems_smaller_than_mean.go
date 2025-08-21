package main

import (
	"fmt"
)

func main() {
	var array = [5]int{}
	var sum float64
	var cnt int
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 5; i++ {
		sum += float64(array[i])
	}
	sum = float64(sum / 5)
	for _, value := range array {
		if float64(value) < sum {
			cnt++
		}
	}
	fmt.Println(cnt)
}
