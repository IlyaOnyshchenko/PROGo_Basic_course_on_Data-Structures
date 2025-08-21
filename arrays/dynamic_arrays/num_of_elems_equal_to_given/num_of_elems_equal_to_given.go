package main

import (
	"fmt"
)

func main() {
	var number, cnt, reqNumber int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&reqNumber)
	for i := 0; i < number; i++ {
		if reqNumber == array[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
