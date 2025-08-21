package main

import (
	"fmt"
)

func main() {
	var array = [5]int{}
	var cnt int = 1
	var flag bool
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 4; i++ {
		if array[i]+array[i+1] == 2 {
			cnt++
			flag = true
		} else if array[i]+array[i+1] == 1 {
			flag = true
		}
	}
	if !flag {
		fmt.Println(0)
	} else {
		fmt.Println(cnt)
	}
}
