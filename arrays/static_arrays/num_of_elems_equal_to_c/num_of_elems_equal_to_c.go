package main

import (
	"fmt"
)

func main() {
	var array = [10]int{}
	var number, cnt int
	for i := 0; i < 10; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&number)
	for i := 0; i < 10; i++ {
		if array[i] == number {
			cnt++
		}
	}
	fmt.Println(cnt)
}
